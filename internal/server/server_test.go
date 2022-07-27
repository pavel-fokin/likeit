package server

import (
	"testing"
)

func TestServerStartAndShutdown(t *testing.T) {
	// setup
	s := New("8080")

	// test
	go s.Start()
	s.Shutdown()
}
