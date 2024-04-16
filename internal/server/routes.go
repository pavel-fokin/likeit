package server

import (
	"io/fs"
	"net/http"
)

type LikeIt interface {
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

func (s *Server) SetupLikesAPIRoutes(likes LikeIt) {
	s.router.Get("/api/likes", getLikes(likes))
	s.router.Post("/api/likes", postLikes(likes))
}
