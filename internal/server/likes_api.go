package server

import (
	"context"
	"log/slog"
	"net/http"
)

type LikesCounter interface {
	CountLikes(ctx context.Context) (int, error)
}

type LikesIncrementor interface {
	IncrementLikes(ctx context.Context) error
}

func getLikes(likes LikesCounter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		count, err := likes.CountLikes(r.Context())
		if err != nil {
			slog.ErrorContext(r.Context(), "failed to get likes", "err", err)
			asErrorResponse(w, err, http.StatusInternalServerError)
			return
		}

		asSuccessResponse(
			w,
			respGetLikes{
				Likes: count,
			},
			http.StatusOK,
		)
	}
}

func postLikes(likes LikesIncrementor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := likes.IncrementLikes(r.Context()); err != nil {
			slog.ErrorContext(r.Context(), "failed to increment likes", "err", err)
			asErrorResponse(w, err, http.StatusInternalServerError)
			return
		}
		asSuccessResponse(w, nil, http.StatusNoContent)
	}
}
