package api

import (
	"context"
	"log/slog"
	"net/http"

	"pavel-fokin/likeit/internal/app"
	"pavel-fokin/likeit/internal/server/apiutil"
)

type Likes interface {
	CountLikes(ctx context.Context) (app.Likes, error)
	IncrementLikes(ctx context.Context) error
}

func GetLikes(app Likes) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		count, err := app.CountLikes(r.Context())
		if err != nil {
			slog.ErrorContext(r.Context(), "failed to get likes", "err", err)
			apiutil.AsErrorResponse(w, err, http.StatusInternalServerError)
			return
		}

		apiutil.AsSuccessResponse(
			w,
			&GetLikesResponse{
				Likes: int(count),
			},
			http.StatusOK,
		)
	}
}

func PostLikes(app Likes) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := app.IncrementLikes(r.Context()); err != nil {
			slog.ErrorContext(r.Context(), "failed to increment likes", "err", err)
			apiutil.AsErrorResponse(w, err, http.StatusInternalServerError)
			return
		}

		apiutil.AsSuccessResponse(w, nil, http.StatusNoContent)
	}
}
