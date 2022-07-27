package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v6"

	"pavel-fokin/likeit/internal/server"
)

var (
	//go:embed web/dist
	web embed.FS
)

type Config struct {
	Server server.Config
}

func ReadConfig() *Config {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	return cfg
}

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	staticFS, _ := fs.Sub(web, "web/dist")

	config := ReadConfig()

	httpServer := server.New(config.Server)
	httpServer.SetupStaticRoutes(staticFS)

	go httpServer.Start()

	<-sig

	httpServer.Shutdown()
}
