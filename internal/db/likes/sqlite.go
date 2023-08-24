package likes

import (
	"context"
	"database/sql"
)

type Likes struct {
	db *sql.DB
}

func New(db *sql.DB) *Likes {
	return &Likes{
		db: db,
	}
}

func (l *Likes) Count(ctx context.Context) (int, error) {
	var count int
	l.db.QueryRow("SELECT count FROM likes;").Scan(&count)
	return count, nil
}

func (l *Likes) Increment(ctx context.Context) error {
	l.db.Exec("UPDATE likes SET count = count + 1;")
	return nil
}
