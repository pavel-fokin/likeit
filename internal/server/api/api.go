package api

import (
	"context"
	"net/http"
	"pavel-fokin/likeit/internal/server/httputil"
)

type LikesCounter interface {
	Count(ctx context.Context) (int, error)
}

type LikesIncrementor interface {
	Increment(ctx context.Context) error
}

func LikesGet(likes LikesCounter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		count, err := likes.Count(r.Context())
		if err != nil {
			httputil.AsErrorResponse(w, err, http.StatusInternalServerError)
			return
		}

		httputil.AsSuccessResponse(
			w,
			LikesResp{
				Likes: count,
			},
			http.StatusOK,
		)
	}
}

func LikesPost(likes LikesIncrementor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    if err := likes.Increment(r.Context()); err != nil {
			httputil.AsErrorResponse(w, err, http.StatusInternalServerError)
			return
    }
		httputil.AsSuccessResponse(w, nil, http.StatusNoContent)
	}
}
