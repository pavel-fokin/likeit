package db

import (
	"database/sql"
	"log"
	"pavel-fokin/likeit/internal/likeit"

	_ "modernc.org/sqlite"
)

type Config struct {
	DATABASE_TYPE string `env:"LIKEIT_DATABASE_TYPE" envDefault:"sqlite"`
	DATABASE_URL  string `env:"LIKEIT_DATABASE_URL" envDefault:":memory:"`
}

type closeFunc func() error

// New creates a new database connection based on the provided configuration.
// It returns an instance of the likeit.DB interface and a close function to release resources.
func New(config Config) (likeit.DB, closeFunc) {
	switch config.DATABASE_TYPE {
	case "sqlite":
		return newSqliteDB(config)
	default:
		log.Fatalf("unsupported database type: %s", config.DATABASE_TYPE)
		return nil, nil
	}
}

func newSqliteDB(config Config) (likeit.DB, closeFunc) {
	conn, err := sql.Open("sqlite", config.DATABASE_URL)
	if err != nil {
		log.Fatal(err)
	}

	// Create initial DB.
	_, err = conn.Exec(SchemaSqlite)
	if err != nil {
		log.Fatal(err)
	}

	return &LikeItSqlite{db: conn}, conn.Close
}
