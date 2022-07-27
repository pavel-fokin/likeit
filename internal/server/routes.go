package server

import (
	"io/fs"
	"net/http"

	"pavel-fokin/likeit/internal/server/api"
)

type Likes interface {
	Count()
	Increment()
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
	s.router.Get("/api/v1/likes", api.LikesGet(likes))
	s.router.Post("/api/v1/likes", api.LikesPost(likes))
}
