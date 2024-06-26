package app

import (
	"context"
)

// Likes is a domain model(type) that represents the number of likes.
type Likes int

// CountLikes returns the number of Likes.
func (a *App) CountLikes(ctx context.Context) (Likes, error) {
	likesCount, err := a.db.CountLikes(ctx)
	if err != nil {
		return 0, err
	}

	return likesCount, nil
}

// IncrementLikes increments the number of Likes.
func (a *App) IncrementLikes(ctx context.Context) error {
	if err := a.db.IncrementLikes(ctx); err != nil {
		return err
	}

	return nil
}
