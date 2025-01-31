// Package db provides the methods to interact with the MySQL/Maria Database used by the auth and character features of a WoW private server.
package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

type Config struct {
	Name   string
	Host   string
	Port   int
	User   string
	Pass   string
	Secure bool
}

type Service struct {
	DB *sql.DB
}

func New(config *Config) *Service {
	s := Service{}
	cfg := mysql.Config{
		User:   config.User,
		Passwd: config.Pass,
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:%d", config.Host, config.Port),
		DBName: config.Name,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	s.DB = db

	return &s
}
