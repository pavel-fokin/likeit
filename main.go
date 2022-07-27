package main

import (
	"embed"
	"io/fs"
	"os"
	"os/signal"
	"syscall"

	"pavel-fokin/likeit/internal/server"
)

const (
	defaultPort = "8080"
)

var (
	//go:embed web/dist
	web embed.FS
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	staticFS, _ := fs.Sub(web, "web/dist")

	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}

	router := server.NewStatic(staticFS)
	httpServer := server.New(port, router)

	go httpServer.Start()

	<-sig

	httpServer.Shutdown()
}
