// Package app contains the application layer of the "LikeIt" service.
package app

import "context"

// Likes is a domain model(type) that represents the number of likes.
type Likes int

// DB is the interface that wraps the basic operations of the database.
type DB interface {
	CountLikes(ctx context.Context) (Likes, error)
	IncrementLikes(ctx context.Context) error
	CreateUser(ctx context.Context) (*User, error)
	FindUser(ctx context.Context, userID string) (*User, error)
}

// App is the application that exposes use cases of the "LikeIt" service.
type App struct {
	db DB
}

// New creates a new instance of the "LikeIt" application.
func New(db DB) *App {
	return &App{
		db: db,
	}
}
