package likeit

import (
	"context"
)

type LikesDB interface {
	Count(ctx context.Context) (int, error)
	Increment(ctx context.Context) error
}

type LikeIt struct {
	likesdb LikesDB
}

func New(db LikesDB) *LikeIt {
	return &LikeIt{
		likesdb: db,
	}
}

func (l *LikeIt) CountLikes(ctx context.Context) (int, error) {
	likesCount, err := l.likesdb.Count(ctx)
	if err != nil {
		return likesCount, err
	}
	return likesCount, nil
}

func (l *LikeIt) IncrementLikes(ctx context.Context) error {
	return l.likesdb.Increment(ctx)
}
