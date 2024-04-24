package server

import (
	"testing"
)

func TestServerStartAndShutdown(t *testing.T) {
	s := New(Config{})
	go s.Start()
	s.Shutdown()
}

func TestSetupLikesAPI(t *testing.T) {
	s := New(Config{})
	s.SetupLikesAPI(nil)
}

func TestSetupAuthAPI(t *testing.T) {
	s := New(Config{})
	s.SetupAuthAPI(nil)
}

func TestSetupStaticRoutes(t *testing.T) {
	s := New(Config{})
	s.SetupStaticRoutes(nil)
}
