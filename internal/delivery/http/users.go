package http

import (
	"net/http"

	"github.com/andriimwks/go-fiber-template/internal/service"
	"github.com/andriimwks/go-fiber-template/pkg/templatetags"
	"github.com/gofiber/fiber/v2"
)

func (h *handler) initUserHandlers() {
	h.app.Get("/", h.indexGET)
	h.app.Get("/signin", h.signInGET)
	h.app.Post("/signin", h.signInPOST)
	h.app.Get("/signup", h.signUpGET)
	h.app.Post("/signup", h.signUpPOST)
	h.app.Get("/signout", h.signOutGET)

	templatetags.AddURL("index", "/")
	templatetags.AddURL("user:signin", "/signin")
	templatetags.AddURL("user:signup", "/signup")
	templatetags.AddURL("user:signout", "/signout")
}

func (h *handler) indexGET(c *fiber.Ctx) error {
	return c.Render("pages/index", fiber.Map{})
}

func (h *handler) signInGET(c *fiber.Ctx) error {
	if c.Locals("user") != nil {
		return c.Redirect(templatetags.GetURL("index"))
	}
	return c.Render("pages/signin", fiber.Map{})
}

func (h *handler) signInPOST(c *fiber.Ctx) error {
	if c.Locals("user") != nil {
		return c.Redirect(templatetags.GetURL("index"))
	}

	var in service.SignInInput
	if err := c.BodyParser(&in); err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	tp, err := h.services.Users.SignIn(in)
	if err != nil {
		return c.Render("pages/signin", fiber.Map{"errors": []string{err.Error()}})
	}

	c.Cookie(&fiber.Cookie{Name: "ACCESS_TOKEN", Value: tp.AccessToken, Secure: true, HTTPOnly: true})
	c.Cookie(&fiber.Cookie{Name: "REFRESH_TOKEN", Value: tp.RefreshToken, Secure: true, HTTPOnly: true})
	return c.Redirect(templatetags.GetURL("index"))
}

func (h *handler) signUpGET(c *fiber.Ctx) error {
	if c.Locals("user") != nil {
		return c.Redirect(templatetags.GetURL("index"))
	}
	return c.Render("pages/signup", fiber.Map{})
}

func (h *handler) signUpPOST(c *fiber.Ctx) error {
	if c.Locals("user") != nil {
		return c.Redirect(templatetags.GetURL("index"))
	}

	var in service.SignUpInput
	if err := c.BodyParser(&in); err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	tp, err := h.services.Users.SignUp(in)
	if err != nil {

		return c.Render("pages/signup", fiber.Map{"errors": []string{err.Error()}})
	}

	c.Cookie(&fiber.Cookie{Name: "ACCESS_TOKEN", Value: tp.AccessToken, Secure: true, HTTPOnly: true})
	c.Cookie(&fiber.Cookie{Name: "REFRESH_TOKEN", Value: tp.RefreshToken, Secure: true, HTTPOnly: true})
	return c.Redirect(templatetags.GetURL("index"))
}

func (h *handler) signOutGET(c *fiber.Ctx) error {
	c.ClearCookie("ACCESS_TOKEN", "REFRESH_TOKEN")
	return c.Redirect(templatetags.GetURL("index"))
}
