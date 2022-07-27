package api

import (
	"net/http"
)

type LikesCounter interface {
	Count()
}

type LikesIncrementor interface {
	Increment()
}

func LikesGet(likes LikesCounter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		likes.Count()
		w.Write([]byte("OK"))
	}
}

func LikesPost(likes LikesIncrementor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		likes.Increment()
		w.Write([]byte("OK"))
	}
}
