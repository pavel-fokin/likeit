package server

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

const (
	_readTimeoutSecs     = 5
	_writeTimeoutSecs    = 5
	_shutdownTimeoutSecs = 5
)

type Server struct {
	server *http.Server
	router chi.Router
}

func New(port string) *Server {
	router := chi.NewRouter()

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  _readTimeoutSecs * time.Second,
		WriteTimeout: _writeTimeoutSecs * time.Second,
	}

	return &Server{
		server: server,
		router: router,
	}
}

func (s *Server) Start() {
	s.server.ListenAndServe()
}

func (s *Server) Shutdown() {
	shutdownCtx, cancelShutdownCtx := context.WithTimeout(
		context.Background(), _shutdownTimeoutSecs*time.Second,
	)
	defer cancelShutdownCtx()

	s.server.Shutdown(shutdownCtx)
}
