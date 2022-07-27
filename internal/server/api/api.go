package api

import (
	"net/http"
)

func LikesGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}
}

func LikesPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}
}
