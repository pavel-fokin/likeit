package likes

import (
	"context"
)

type LikesDB interface {
	Count(ctx context.Context) (int, error)
	Increment(ctx context.Context) error
}

type Likes struct {
	db LikesDB
}

func New(db LikesDB) *Likes {
	return &Likes{
		db: db,
	}
}

func (l *Likes) Count(ctx context.Context) (int, error) {
	likesCount, err := l.db.Count(ctx)
	if err != nil {
		return likesCount, err
	}
	return likesCount, nil
}

func (l *Likes) Increment(ctx context.Context) error {
	return l.db.Increment(ctx)
}
