package cmd

import (
	"context"

	"github.com/softwr-skullclown/azeroth-registration/internal/http"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Launches the example webapp on https://localhost:8080",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}

func serve() {
	h := http.New(http.Config{
		ListenAddress: c.ListenAddress,
	})

	err := h.ListenAndServe(context.Background())
	if err != nil {
		panic(err)
	}
}
