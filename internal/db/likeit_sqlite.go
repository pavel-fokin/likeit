// This package contains the implementation of the LikeItSqlite struct.
package db

import (
	"context"
	"database/sql"
	"fmt"

	"pavel-fokin/likeit/internal/app"
)

type LikeItSqlite struct {
	db *sql.DB
}

var _ app.DB = (*LikeItSqlite)(nil)

// NewLikesSqlite creates a new LikesSqlite instance.
func NewLikeItSqlite(db *sql.DB) *LikeItSqlite {
	return &LikeItSqlite{
		db: db,
	}
}

// CountLikes returns the number of likes.
func (l *LikeItSqlite) CountLikes(ctx context.Context) (app.Likes, error) {
	var count int
	err := l.db.QueryRowContext(ctx, "SELECT count FROM likes;").Scan(&count)
	if err != nil {
		return app.Likes(0), fmt.Errorf("failed to select likes count: %w", err)
	}
	return app.Likes(count), nil
}

// IncrementLikes increments the number of likes.
func (l *LikeItSqlite) IncrementLikes(ctx context.Context) error {
	_, err := l.db.ExecContext(ctx, "UPDATE likes SET count = count + 1;")
	if err != nil {
		return fmt.Errorf("failed to update likes with increment: %w", err)
	}
	return nil
}
