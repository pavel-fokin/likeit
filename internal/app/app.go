package app

import "context"

// DB is the interface that wraps the basic operations of the database.
type DB interface {
	CountLikes(ctx context.Context) (Likes, error)
	IncrementLikes(ctx context.Context) error
}

// App is the application that exposes use cases of "LikeIt".
type App struct {
	db DB
}

func New(db DB) *App {
	return &App{
		db: db,
	}
}
