package main

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v6"

	"pavel-fokin/likeit/internal/db"
	"pavel-fokin/likeit/internal/likeit"
	"pavel-fokin/likeit/internal/server"
	"pavel-fokin/likeit/web"
)

type Config struct {
	Server server.Config
	DB     db.Config
}

func readConfig() *Config {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	return cfg
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	staticFS, _ := fs.Sub(web.Dist, "dist")

	config := readConfig()

	likeItDB, close := db.New(config.DB)
	defer close()

	likeIt := likeit.New(likeItDB)

	httpServer := server.New(ctx, config.Server)
	httpServer.SetupStaticRoutes(staticFS)
	httpServer.SetupLikesAPIRoutes(likeIt)

	go httpServer.Start()

	<-ctx.Done()

	if err := httpServer.Shutdown(); err != nil {
		log.Fatal("Failed to shutdown the server gracefully")
	}
}
