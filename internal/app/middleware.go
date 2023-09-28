package app

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	loggerMiddleware = logger.New(logger.Config{
		Format:       "[${time}] ${ip}:${port} ${method} ${path} ${status} (${latency})\n",
		TimeFormat:   time.RFC1123,
		TimeZone:     "Local",
		TimeInterval: 1 * time.Second,
	})
	compressMiddleware = compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	})
)

func corsMiddleware(origins, methods string) func(*fiber.Ctx) error {
	return cors.New(cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     methods,
		AllowCredentials: true,
	})
}

func csrfMiddleware(exp time.Duration) func(*fiber.Ctx) error {
	return csrf.New(csrf.Config{
		Next:           func(_ *fiber.Ctx) bool { return false },
		KeyLookup:      "cookie:CSRF_TOKEN",
		Expiration:     exp,
		CookieName:     "CSRF_TOKEN",
		CookieSecure:   true,
		CookieHTTPOnly: true,
	})
}

func limiterMiddleware(max int, exp time.Duration) func(*fiber.Ctx) error {
	return limiter.New(limiter.Config{
		Max:        max,
		Expiration: exp,
	})
}
