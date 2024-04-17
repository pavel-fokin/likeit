// Package app contains the application layer of the "LikeIt" service.
package app

import "context"

// DB is the interface that wraps the basic operations of the database.
type DB interface {
	CountLikes(ctx context.Context) (Likes, error)
	IncrementLikes(ctx context.Context) error
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
