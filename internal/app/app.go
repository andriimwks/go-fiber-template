package app

import (
	"github.com/andriimwks/go-fiber-template/internal/config"
	"github.com/andriimwks/go-fiber-template/internal/delivery/http"
	"github.com/andriimwks/go-fiber-template/internal/repository"
	"github.com/andriimwks/go-fiber-template/internal/service"
	"github.com/gofiber/fiber/v2"
)

func New(cfg *config.Config) (*fiber.App, error) {
	r := newRouter(cfg)

	db, err := newDatabase(cfg)
	if err != nil {
		return nil, err
	}

	http.Init(r, cfg, service.New(
		cfg,
		repository.New(db),
	))

	return r, nil
}
