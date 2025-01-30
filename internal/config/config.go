package config

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/spf13/viper"
)

func New(ctx context.Context, cfgFile string) *Config {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".webapp-example") // name of config file (without extension)
	viper.AddConfigPath("$HOME")           // adding home directory as first search path

	viper.SetDefault("LISTEN_ADDRESS", ":8080")
	viper.SetDefault("GAME_VERSION", "3.3.5a") // default to WoTLK most popular patch version
	viper.SetDefault("PAGE_TITLE", "My Private Server Registration")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			slog.WarnContext(ctx, fmt.Sprintf("Config file not found; ignore error if desired: %v", err))
		} else {
			slog.ErrorContext(ctx, fmt.Sprintf("Config file was found but another error was produced: %v", err))
		}
	}

	config := Config{}
	err := viper.Unmarshal(&config)
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("unable to unmarshal config into struct, %v", err))
	}

	return &config
}
