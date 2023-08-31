package likes

import (
	"context"
	"database/sql"
	"fmt"
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
	err := l.db.QueryRow("SELECT count FROM likes;").Scan(&count)
	if err != nil {
		return count, fmt.Errorf("failed to select likes: %w", err)
	}
	return count, nil
}

func (l *Likes) Increment(ctx context.Context) error {
	_, err := l.db.Exec("UPDATE likes SET count = count + 1;")
	if err != nil {
		return fmt.Errorf("failed to update likes: %w", err)
	}
	return nil
}
