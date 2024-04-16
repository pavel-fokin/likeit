package server

import (
	"log/slog"
	"net/http"
)

func getLikes(likes LikesCounter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		count, err := likes.CountLikes(r.Context())
		if err != nil {
			slog.ErrorContext(r.Context(), "failed to get likes", err)
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
			slog.ErrorContext(r.Context(), "failed to increment likes", err)
			asErrorResponse(w, err, http.StatusInternalServerError)
			return
		}
		asSuccessResponse(w, nil, http.StatusNoContent)
	}
}
