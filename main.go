package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
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

	log.Fatal(http.ListenAndServe(":8080", nil))
}
