package server

import (
	"context"
	"testing"
	"time"
)

func TestServerStartAndShutdown(t *testing.T) {
	t.Run("Start and shutdown", func(t *testing.T) {
		s := New(context.Background(), Config{})

		go s.Start()
		s.Shutdown()
	})

	t.Run("Start and shutdown with context canceled", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		s := New(ctx, Config{})

		go s.Start()

		time.Sleep(50 * time.Millisecond)
		cancel()
	})
}

func TestServerSetupLikesAPIRoutes(t *testing.T) {
	s := New(context.Background(), Config{})

	s.SetupLikesAPIRoutes(nil)
}

func TestServerSetupStaticRoutes(t *testing.T) {
	s := New(context.Background(), Config{})

	s.SetupStaticRoutes(nil)
}
