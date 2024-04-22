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

func (s *Server) SetupAPIRoutes(likes api.LikeIt) {
	s.router.Get("/api/likes", api.GetLikes(likes))
	s.router.Post("/api/likes", api.PostLikes(likes))
}

func (s *Server) SetupAuthRoutes(auth api.AppAuth) {
	s.router.Post("/api/auth/signin", api.SignIn(auth, s.config.tokenSigningKey))
	s.router.Post("/api/auth/signup", api.SignUp(auth, s.config.tokenSigningKey))
}
