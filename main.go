package main

import (
	"embed"
	"os"
	"io/fs"
	"log"
	"net/http"
)

const (
	defaultPort = "8080"
)

var (
	//go:embed dist
	dist embed.FS
)

func main() {
	web, _ := fs.Sub(dist, "dist")

	http.Handle(
		"/", http.FileServer(http.FS(web)),
	)

	port := os.Getenv("PORT")

	if port == "" {
  	port = defaultPort
	}

	log.Fatal(http.ListenAndServe(":" + port, nil))
}
