package middlewares

import (
	"time"

	"github.com/AkifhanIlgaz/taxihub/api-gateway/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func NewRateLimiter(cfg config.RateLimitConfig) fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        cfg.MaxRequests,
		Expiration: time.Duration(cfg.WindowMinutes) * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			userID := c.Locals("userId")
			if userID != nil {
				return userID.(string)
			}
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"success": false,
				"error":   "Rate limit exceeded",
			})
		},
	})
}
