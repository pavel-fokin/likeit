package server

import (
	"testing"
)

func TestServerStartAndShutdown(t *testing.T) {
	s := New(Config{})
	go s.Start()
	s.Shutdown()
}

func TestServerSetupLikesAPIRoutes(t *testing.T) {
	s := New(Config{})
	s.SetupLikesAPIRoutes(nil)
}

func TestServerSetupStaticRoutes(t *testing.T) {
	s := New(Config{})
	s.SetupStaticRoutes(nil)
}
