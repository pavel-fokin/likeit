package server

import (
	"context"
	"io/fs"
	"net/http"
)

type LikesCounter interface {
	CountLikes(ctx context.Context) (int, error)
}

type LikesIncrementor interface {
	IncrementLikes(ctx context.Context) error
}

type Likes interface {
	LikesCounter
	LikesIncrementor
}

func (s *Server) SetupStaticRoutes(static fs.FS) {
	s.router.Handle(
		"/", http.FileServer(http.FS(static)),
	)
	s.router.Handle(
		"/static/*", http.StripPrefix("/static/", http.FileServer(http.FS(static))),
	)
}

func (s *Server) SetupLikesAPIRoutes(likes Likes) {
	s.router.Get("/api/likes", getLikes(likes))
	s.router.Post("/api/likes", postLikes(likes))
}
