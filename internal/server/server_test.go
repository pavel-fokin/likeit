package server

import (
	"testing"
)

func TestServerStartAndShutdown(t *testing.T) {
	// setup
	s := New(Config{})

	// test
	go s.Start()
	s.Shutdown()
}
