package db

import (
	"context"
	"database/sql"
	"fmt"

	"pavel-fokin/likeit/internal/likeit"
)

type LikeItSqlite struct {
	db *sql.DB
}

var _ likeit.DB = (*LikeItSqlite)(nil)

// NewLikesSqlite creates a new LikesSqlite instance.
func NewLikesSqlite(db *sql.DB) *LikeItSqlite {
	return &LikeItSqlite{
		db: db,
	}
}

// CountLikes returns the number of likes.
func (l *LikeItSqlite) CountLikes(ctx context.Context) (int, error) {
	var count int
	err := l.db.QueryRow("SELECT count FROM likes;").Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to select likes count: %w", err)
	}
	return count, nil
}

// IncrementLikes increments the number of likes.
func (l *LikeItSqlite) IncrementLikes(ctx context.Context) error {
	_, err := l.db.Exec("UPDATE likes SET count = count + 1;")
	if err != nil {
		return fmt.Errorf("failed to update likes with increment: %w", err)
	}
	return nil
}
