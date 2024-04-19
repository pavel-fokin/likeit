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

func readConfig() *Config {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		log.Printf("%+v\n", err)
	}
	return cfg
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	staticFS, _ := fs.Sub(web.Dist, "dist")

	config := readConfig()

	likeitDB, close := db.New(config.DB)
	defer close()

	likeitApp := app.New(likeitDB)

	likeitServer := server.New(config.Server)
	likeitServer.SetupStaticRoutes(staticFS)
	likeitServer.SetupLikesAPIRoutes(likeitApp)

	go likeitServer.Start()

	<-ctx.Done()

	if err := likeitServer.Shutdown(); err != nil {
		log.Fatal("Failed to shutdown the server gracefully")
	}
}
