package api

import (
	"fmt"
	"net/http"
)

type LikesCounter interface {
	Count() (int, error)
}

type LikesIncrementor interface {
	Increment() error
}

func LikesGet(likes LikesCounter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		count, _ := likes.Count()
		w.Write([]byte(fmt.Sprintf("%d", count)))
	}
}

func LikesPost(likes LikesIncrementor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		likes.Increment()
		w.Write([]byte("OK"))
	}
}
