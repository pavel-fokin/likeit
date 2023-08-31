package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	DATABASE_URL string `env:"DATABASE_URL" envDefault:":memory:"`
}

func NewSqlite(config Config) *sql.DB {

	conn, err := sql.Open("sqlite3", config.DATABASE_URL)
	if err != nil {
		log.Fatal(err)
	}

	// Create initial DB.
	_, err = conn.Exec(Schema)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}
