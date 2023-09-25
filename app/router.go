package app

import (
	"github.com/andriimwks/go-fiber-template/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/django/v3"
)

func newRouter(cfg *config.Config) *fiber.App {
	r := fiber.New(fiber.Config{
		UnescapePath: true,
		Views:        django.New("./www/templates", ".html"),
		AppName:      "go-fiber-template",
	})
	r.Static("/static", "./www/static")
	r.Use(
		recover.New(),
		loggerMiddleware,
		compressMiddleware,
		corsMiddleware(cfg.Security.CorsAllowOrigins, cfg.Security.CorsAllowMethods),
		csrfMiddleware(cfg.Security.CsrfExpiration),
		limiterMiddleware(cfg.Limiter.Max, cfg.Limiter.Expiration),
	)
	return r
}
