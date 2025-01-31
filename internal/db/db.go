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

func New(c *Config) *Service {
	s := Service{}
	dbaddr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	cfg := mysql.Config{
		User:   c.User,
		Passwd: c.Pass,
		Addr:   dbaddr,
		DBName: c.Name,
		Net:    "tcp",
	}
	dsn := cfg.FormatDSN()

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	s.DB = db

	return &s
}
