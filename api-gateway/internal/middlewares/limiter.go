package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func NewRateLimiter(maxRequests int, duration time.Duration) fiber.Handler {
	config := limiter.Config{
		Max:        maxRequests,
		Expiration: duration,
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
	}

	return limiter.New(config)
}
