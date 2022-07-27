package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	URL string `env:"DATABASE_URL" envDefault:"db.sqlite3"`
}

func New(config Config) *sql.DB {

	conn, err := sql.Open("sqlite3", config.URL)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}
