package main

import (
	"fmt"
	"io/fs"
	"os"
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
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	staticFS, _ := fs.Sub(web.Dist, "dist")

	config := readConfig()

	d := db.NewSqliteDB(config.DB)
	defer d.Close()

	likesDB := db.NewLikesSqlite(d)

	likeIt := likeit.New(likesDB)

	httpServer := server.New(config.Server)
	httpServer.SetupStaticRoutes(staticFS)
	httpServer.SetupLikesAPIRoutes(likeIt)

	go httpServer.Start()

	<-sig

	httpServer.Shutdown()
}
