package app

import (
	"context"
)

// CountLikes returns the number of Likes.
func (l *App) CountLikes(ctx context.Context) (Likes, error) {
	likesCount, err := l.db.CountLikes(ctx)
	if err != nil {
		return 0, err
	}
	return likesCount, nil
}

// IncrementLikes increments the number of Likes.
func (l *App) IncrementLikes(ctx context.Context) error {
	if err := l.db.IncrementLikes(ctx); err != nil {
		return err
	}
	return nil
}
