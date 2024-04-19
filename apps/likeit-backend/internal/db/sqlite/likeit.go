// This package contains the implementation of the LikeItSqlite struct.
package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"

	"pavel-fokin/likeit/internal/app"
	"pavel-fokin/likeit/internal/db/dbutil"
)

// LikeIt is a SQLite implementation of the app.DB interface.
type LikeIt struct {
	db *sql.DB
}

var _ app.DB = (*LikeIt)(nil)

func New(url string) (app.DB, func() error) {
	db, err := sql.Open("sqlite", url)
	if err != nil {
		log.Fatal(err)
	}

	// Create initial DB.
	_, err = db.Exec(SchemaSqlite)
	if err != nil {
		log.Fatal(err)
	}

	return &LikeIt{
		db: db,
	}, db.Close
}

// CountLikes returns the number of likes.
func (l *LikeIt) CountLikes(ctx context.Context) (app.Likes, error) {
	var count int
	err := l.db.QueryRowContext(ctx, "SELECT count FROM likes;").Scan(&count)
	if err != nil {
		return app.Likes(0), fmt.Errorf("failed to select likes count: %w", err)
	}
	return app.Likes(count), nil
}

// IncrementLikes increments the number of likes.
func (l *LikeIt) IncrementLikes(ctx context.Context) error {
	_, err := l.db.ExecContext(ctx, "UPDATE likes SET count = count + 1;")
	if err != nil {
		return fmt.Errorf("failed to update likes with increment: %w", err)
	}
	return nil
}

// CreateUser creates a new user.
func (l *LikeIt) CreateUser(ctx context.Context) (*app.User, error) {
	id, err := dbutil.RandomID()
	if err != nil {
		return nil, fmt.Errorf("failed to generate random ID: %w", err)
	}

	_, err = l.db.ExecContext(ctx, "INSERT INTO users (id) VALUES (?);", id)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}

	return &app.User{
		ID: id,
	}, nil
}

// FindUser finds a user by ID.
func (l *LikeIt) FindUser(ctx context.Context, id string) (*app.User, error) {
	var userID string
	err := l.db.QueryRowContext(ctx, "SELECT id FROM users WHERE id = ?;", id).Scan(&userID)
	if err != nil {
		return nil, fmt.Errorf("failed to select user: %w", err)
	}

	return &app.User{
		ID: userID,
	}, nil
}
