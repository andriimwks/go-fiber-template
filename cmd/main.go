package main

import (
	"log"

	"github.com/andriimwks/go-fiber-template/internal/app"
	"github.com/andriimwks/go-fiber-template/internal/config"
	_ "github.com/andriimwks/go-fiber-template/pkg/templatetags"
)

func main() {
	cfg, err := config.Load(".", "config")
	if err != nil {
		log.Fatal(err)
	}

	r, err := app.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(r.Listen(cfg.Server.Address))
}
