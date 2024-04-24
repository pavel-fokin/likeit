package server

import (
	"io/fs"
	"net/http"

	"pavel-fokin/likeit/internal/server/api"
)

func (s *Server) SetupStaticRoutes(static fs.FS) {
	fs := http.FileServerFS(static)

	s.router.Get("/", fs.ServeHTTP)
	s.router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/"
		fs.ServeHTTP(w, r)
	})
	s.router.Get(
		"/assets/*", fs.ServeHTTP,
	)
}

func (s *Server) SetupLikesAPI(likes api.LikeIt) {
	s.router.Get("/api/likes", api.GetLikes(likes))
	s.router.Post("/api/likes", api.PostLikes(likes))
}

func (s *Server) SetupAuthAPI(auth api.AppAuth) {
	s.router.Post("/api/auth/signin", api.SignIn(auth, s.config.tokenSigningKey))
	s.router.Post("/api/auth/signup", api.SignUp(auth, s.config.tokenSigningKey))
}
