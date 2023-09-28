package main

import (
	"flag"
	"log"

	"github.com/andriimwks/go-fiber-template/internal/app"
	"github.com/andriimwks/go-fiber-template/internal/config"
	_ "github.com/andriimwks/go-fiber-template/pkg/templatetags"
)

func main() {
	configName := flag.String("config", "dev", "")
	flag.Parse()

	cfg, err := config.Load("./configs", *configName)
	if err != nil {
		log.Fatal(err)
	}

	r, err := app.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(r.Listen(cfg.Server.Address))
}
