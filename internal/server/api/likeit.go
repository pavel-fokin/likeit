package api

import (
	"context"
	"log/slog"
	"net/http"

	"pavel-fokin/likeit/internal/app"
	"pavel-fokin/likeit/internal/server/httputil"
)

type LikesCounter interface {
	CountLikes(ctx context.Context) (app.Likes, error)
}

type LikesIncrementor interface {
	IncrementLikes(ctx context.Context) error
}

type LikeIt interface {
	LikesCounter
	LikesIncrementor
}

func GetLikes(likes LikesCounter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		count, err := likes.CountLikes(r.Context())
		if err != nil {
			slog.ErrorContext(r.Context(), "failed to get likes", "err", err)
			httputil.AsErrorResponse(w, err, http.StatusInternalServerError)
			return
		}

		httputil.AsSuccessResponse(
			w,
			GetLikesResponse{
				Likes: int(count),
			},
			http.StatusOK,
		)
	}
}

func PostLikes(likes LikesIncrementor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := likes.IncrementLikes(r.Context()); err != nil {
			slog.ErrorContext(r.Context(), "failed to increment likes", "err", err)
			httputil.AsErrorResponse(w, err, http.StatusInternalServerError)
			return
		}

		httputil.AsSuccessResponse(w, nil, http.StatusNoContent)
	}
}
