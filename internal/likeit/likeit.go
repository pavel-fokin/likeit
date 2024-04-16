package likeit

import (
	"context"
)

type DB interface {
	CountLikes(ctx context.Context) (int, error)
	IncrementLikes(ctx context.Context) error
}

type LikeIt struct {
	db DB
}

func New(db DB) *LikeIt {
	return &LikeIt{
		db: db,
	}
}

func (l *LikeIt) CountLikes(ctx context.Context) (int, error) {
	likesCount, err := l.db.CountLikes(ctx)
	if err != nil {
		return 0, err
	}
	return likesCount, nil
}

func (l *LikeIt) IncrementLikes(ctx context.Context) error {
	if err := l.db.IncrementLikes(ctx); err != nil {
		return err
	}
	return nil
}
