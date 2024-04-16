package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

type Config struct {
	DATABASE_URL string `env:"DATABASE_URL" envDefault:":memory:"`
}

func NewSqliteDB(config Config) *sql.DB {

	conn, err := sql.Open("sqlite", config.DATABASE_URL)
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
