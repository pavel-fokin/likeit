package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

var (
	//go:embed web
	dist embed.FS
)

func main() {
	web, _ := fs.Sub(fs.FS(dist), "web")

	http.Handle(
		"/", http.FileServer(http.FS(web)),
	)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
