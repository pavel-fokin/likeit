package server

import (
	"context"
	"io/fs"
	"net/http"
	"time"
)

const _shutdownTimeoutSecs = 5

type Server struct {
	server *http.Server
}

func New(port string, static fs.FS) *Server {
	mux := http.NewServeMux()
	mux.Handle(
		"/", http.FileServer(http.FS(static)),
	)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	return &Server{
		server: server,
	}
}

func (s *Server) Run() {
	s.server.ListenAndServe()
}

func (s *Server) Shutdown() {
	shutdownCtx, cancelShutdownCtx := context.WithTimeout(
		context.Background(), _shutdownTimeoutSecs*time.Second,
	)
	defer cancelShutdownCtx()

	s.server.Shutdown(shutdownCtx)
}
