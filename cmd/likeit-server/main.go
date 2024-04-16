package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v6"

	"pavel-fokin/likeit/internal/db"
	likesdb "pavel-fokin/likeit/internal/db/likes"
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

	db := db.NewSqlite(config.DB)
	defer db.Close()

	likesDB := likesdb.New(db)

	likeIt := likeit.New(likesDB)

	httpServer := server.New(config.Server)
	httpServer.SetupStaticRoutes(staticFS)
	httpServer.SetupLikesAPIRoutes(likeIt)

	go httpServer.Start()

	<-sig

	httpServer.Shutdown()
}
