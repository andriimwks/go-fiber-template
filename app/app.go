package app

import (
	"github.com/andriimwks/go-fiber-template/config"
	userHttpDelivery "github.com/andriimwks/go-fiber-template/user/delivery/http"
	userRepo "github.com/andriimwks/go-fiber-template/user/repository/gorm"
	userService "github.com/andriimwks/go-fiber-template/user/service"
	_ "github.com/andriimwks/go-fiber-template/www/templatetags"
	"github.com/gofiber/fiber/v2"
)

func New(cfg *config.Config) (*fiber.App, error) {
	r := newRouter(cfg)

	db, err := newDatabase(cfg)
	if err != nil {
		return nil, err
	}

	// repositories
	userRepo := userRepo.New(db)

	// services
	userService := userService.New(cfg, userRepo)

	// handlers
	userHttpDelivery.Init(r, cfg, userService)

	return r, nil
}
