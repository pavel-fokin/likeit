package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
)

const (
	defaultPort = "8080"
)

var (
	//go:embed web/dist
	web embed.FS
)

func main() {
	dist, _ := fs.Sub(web, "web/dist")

	http.Handle(
		"/", http.FileServer(http.FS(dist)),
	)

	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
