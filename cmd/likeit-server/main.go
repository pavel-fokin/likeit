package main

import (
	"context"
	"fmt"
	"io/fs"
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

	d := db.NewSqliteDB(config.DB)
	defer d.Close()

	likesDB := db.NewLikesSqlite(d)

	likeIt := likeit.New(likesDB)

	httpServer := server.New(ctx, config.Server)
	httpServer.SetupStaticRoutes(staticFS)
	httpServer.SetupLikesAPIRoutes(likeIt)

	go httpServer.Start()

	<-ctx.Done()

	httpServer.Shutdown()
}
