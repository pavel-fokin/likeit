package db

import (
	"context"
	"database/sql"
	"fmt"

	"pavel-fokin/likeit/internal/likeit"
)

type LikesSqlite struct {
	db *sql.DB
}

var _ likeit.LikesDB = &LikesSqlite{}

func NewLikesSqlite(db *sql.DB) *LikesSqlite {
	return &LikesSqlite{
		db: db,
	}
}

func (l *LikesSqlite) Count(ctx context.Context) (int, error) {
	var count int
	err := l.db.QueryRow("SELECT count FROM likes;").Scan(&count)
	if err != nil {
		return count, fmt.Errorf("failed to select likes: %w", err)
	}
	return count, nil
}

func (l *LikesSqlite) Increment(ctx context.Context) error {
	_, err := l.db.Exec("UPDATE likes SET count = count + 1;")
	if err != nil {
		return fmt.Errorf("failed to update likes: %w", err)
	}
	return nil
}
