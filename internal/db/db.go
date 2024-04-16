// Package db contains the database specific code of the "LikeIt" service.
package db

import (
	"log"

	"pavel-fokin/likeit/internal/app"
	"pavel-fokin/likeit/internal/db/sqlite"
)

type Config struct {
	DATABASE_TYPE string `env:"LIKEIT_DATABASE_TYPE" envDefault:"sqlite"`
	DATABASE_URL  string `env:"LIKEIT_DATABASE_URL" envDefault:":memory:"`
}

type closeFunc func() error

// New creates a new database connection based on the provided configuration.
// It returns an instance of the likeit.DB interface and a close function to release resources.
func New(config Config) (app.DB, closeFunc) {
	switch config.DATABASE_TYPE {
	case "sqlite":
		return sqlite.New(config.DATABASE_URL)
	default:
		log.Fatalf("unsupported database type: %s", config.DATABASE_TYPE)
		return nil, nil
	}
}
