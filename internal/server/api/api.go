package api

import (
	"context"
	"fmt"
	"net/http"
)

type LikesCounter interface {
	Count(ctx context.Context) (int, error)
}

type LikesIncrementor interface {
	Increment(ctx context.Context) error
}

func LikesGet(likes LikesCounter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		count, _ := likes.Count(r.Context())
		w.Write([]byte(fmt.Sprintf("%d", count)))
	}
}

func LikesPost(likes LikesIncrementor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		likes.Increment(r.Context())
		w.Write([]byte("OK"))
	}
}
