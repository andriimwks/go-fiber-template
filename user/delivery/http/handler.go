package http

import (
	"net/http"

	"github.com/andriimwks/go-fiber-template/config"
	"github.com/andriimwks/go-fiber-template/user"
	"github.com/andriimwks/go-fiber-template/www/templatetags"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	app     *fiber.App
	config  *config.Config
	service user.Service
}

func Init(app *fiber.App, cfg *config.Config, service user.Service) {
	h := handler{app, cfg, service}
	h.app.Use(h.authMiddleware)

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

func (handler) indexGET(c *fiber.Ctx) error {
	return c.Render("pages/index", fiber.Map{})
}

func (handler) signInGET(c *fiber.Ctx) error {
	if c.Locals("user") != nil {
		return c.Redirect(templatetags.GetURL("index"))
	}
	return c.Render("pages/signin", fiber.Map{})
}

func (h *handler) signInPOST(c *fiber.Ctx) error {
	if c.Locals("user") != nil {
		return c.Redirect(templatetags.GetURL("index"))
	}

	var form user.SignInForm
	if err := c.BodyParser(&form); err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	tp, err := h.service.SignIn(form)
	if err != nil {
		var msg string

		switch err {
		case user.ErrInvalidEmail:
			msg = "Invalid email address"
		case user.ErrIncorrectEmailOrPassword:
			msg = "Invalid email and/or password"
		}

		return c.Render("pages/signin", fiber.Map{"errors": []string{msg}})
	}

	c.Cookie(&fiber.Cookie{Name: "ACCESS_TOKEN", Value: tp.AccessToken, Secure: true, HTTPOnly: true})
	c.Cookie(&fiber.Cookie{Name: "REFRESH_TOKEN", Value: tp.RefreshToken, Secure: true, HTTPOnly: true})
	return c.Redirect(templatetags.GetURL("index"))
}

func (handler) signUpGET(c *fiber.Ctx) error {
	if c.Locals("user") != nil {
		return c.Redirect(templatetags.GetURL("index"))
	}
	return c.Render("pages/signup", fiber.Map{})
}

func (h *handler) signUpPOST(c *fiber.Ctx) error {
	if c.Locals("user") != nil {
		return c.Redirect(templatetags.GetURL("index"))
	}

	var form user.SignUpForm
	if err := c.BodyParser(&form); err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	tp, err := h.service.SignUp(form)
	if err != nil {
		var msg string

		switch err {
		case user.ErrEmptyFirstName:
			msg = "Empty first name"
		case user.ErrInvalidNickname:
			msg = "Invalid nickname"
		case user.ErrInvalidEmail:
			msg = "Invalid email address"
		case user.ErrInvalidPassword:
			msg = "Weak password"
		case user.ErrUserExists:
			msg = "Such user already exists"
		}

		return c.Render("pages/signup", fiber.Map{"errors": []string{msg}})
	}

	c.Cookie(&fiber.Cookie{Name: "ACCESS_TOKEN", Value: tp.AccessToken, Secure: true, HTTPOnly: true})
	c.Cookie(&fiber.Cookie{Name: "REFRESH_TOKEN", Value: tp.RefreshToken, Secure: true, HTTPOnly: true})
	return c.Redirect(templatetags.GetURL("index"))
}

func (handler) signOutGET(c *fiber.Ctx) error {
	c.ClearCookie("ACCESS_TOKEN", "REFRESH_TOKEN")
	return c.Redirect(templatetags.GetURL("index"))
}
