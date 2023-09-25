package main

import (
	"flag"
	"log"

	"github.com/andriimwks/go-fiber-template/app"
	"github.com/andriimwks/go-fiber-template/config"
)

func main() {
	configName := flag.String("config", "dev", "")
	flag.Parse()

	cfg, err := config.Load("./config", *configName)
	if err != nil {
		log.Fatal(err)
	}

	r, err := app.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(r.Listen(cfg.Server.Address))
}
