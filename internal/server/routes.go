package server

import (
	"io/fs"
	"net/http"

	"pavel-fokin/likeit/internal/server/api"
)

func (s *Server) SetupStaticRoutes(static fs.FS) {
	s.router.Handle(
		"/", http.FileServer(http.FS(static)),
	)
	s.router.Handle(
		"/static/*", http.StripPrefix("/static/", http.FileServer(http.FS(static))),
	)
}

func (s *Server) SetupLikesAPIRoutes() {
	s.router.Get("/api/v1/likes", api.LikesGet())
	s.router.Post("/api/v1/likes", api.LikesPost())
}
