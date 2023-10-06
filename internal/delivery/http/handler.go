package http

import (
	"github.com/andriimwks/go-fiber-template/internal/config"
	"github.com/andriimwks/go-fiber-template/internal/service"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	app      *fiber.App
	config   *config.Config
	services *service.Services
}

func Init(app *fiber.App, cfg *config.Config, services *service.Services) {
	h := handler{app, cfg, services}
	h.app.Use(h.authMiddleware)
	h.initUserHandlers()
}
