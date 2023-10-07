package app

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	loggerMiddleware = logger.New(logger.Config{
		Format:       "[${time}] ${ip}:${port} ${method} ${path} ${status} (${latency})\n",
		TimeFormat:   time.RFC1123,
		TimeZone:     "Local",
		TimeInterval: 1 * time.Second,
	})
	csrfMiddleware = csrf.New(csrf.Config{
		Next:           func(_ *fiber.Ctx) bool { return false },
		KeyLookup:      "cookie:CSRF_TOKEN",
		CookieName:     "CSRF_TOKEN",
		CookieSecure:   true,
		CookieHTTPOnly: true,
		Expiration:     1 * time.Hour,
	})
)
