package http

import (
	"github.com/andriimwks/go-fiber-template/internal/service"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) authMiddleware(c *fiber.Ctx) error {
	tp, err := h.services.Users.ValidateTokens(service.TokenPair{
		AccessToken:  c.Cookies("ACCESS_TOKEN"),
		RefreshToken: c.Cookies("REFRESH_TOKEN"),
	})
	if err != nil {
		c.ClearCookie("ACCESS_TOKEN", "REFRESH_TOKEN")
		return c.Next()
	}

	if tp.MustUpdate {
		c.Cookie(&fiber.Cookie{Name: "ACCESS_TOKEN", Value: tp.AccessToken, Secure: true, HTTPOnly: true})
		c.Cookie(&fiber.Cookie{Name: "REFRESH_TOKEN", Value: tp.RefreshToken, Secure: true, HTTPOnly: true})
	}

	c.Locals("user", tp.Subject)
	c.Bind(fiber.Map{"user": fiber.Map{"id": tp.Subject.ID}})
	return c.Next()
}
