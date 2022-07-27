package server

import (
	"io/fs"
	"net/http"
)

func (s *Server) SetupStaticRoutes(static fs.FS) {
	s.router.Handle(
		"/", http.FileServer(http.FS(static)),
	)
	s.router.Handle(
		"/static/*", http.StripPrefix("/static/",http.FileServer(http.FS(static))),
	)
}
