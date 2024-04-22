package main

import (
	"context"
	"io/fs"
	"log"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v6"

	"pavel-fokin/likeit/internal/app"
	"pavel-fokin/likeit/internal/db"
	"pavel-fokin/likeit/internal/server"
	"pavel-fokin/likeit/web"
)

type Config struct {
	Server server.Config
	DB     db.Config
}

func NewConfig() *Config {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		log.Printf("%+v\n", err)
	}
	return cfg
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	config := NewConfig()

	likeitDB, close := db.New(config.DB)
	defer close()

	likeitApp := app.New(likeitDB)

	likeitServer := server.New(config.Server)
	// Setup static routes.
	staticFS, _ := fs.Sub(web.Dist, "dist")
	likeitServer.SetupStaticRoutes(staticFS)
	// Setup API routes.
	likeitServer.SetupAuthRoutes(likeitApp)
	likeitServer.SetupAPIRoutes(likeitApp)

	log.Println("Starting LikeIt HTTP server... ", config.Server.Port)
	go likeitServer.Start()

	<-ctx.Done()

	log.Println("Shutting down the LikeIt HTTP server...")
	if err := likeitServer.Shutdown(); err != nil {
		log.Fatal("Failed to shutdown the server gracefully")
	}
	log.Println("LikeIt HTTP server shutdown successfully")
}
