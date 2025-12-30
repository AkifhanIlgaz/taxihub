package middlewares

import (
	"strings"

	"github.com/AkifhanIlgaz/taxihub/api-gateway/pkg/token"
	"github.com/gofiber/fiber/v2"
)

const bearerPrefix = "Bearer "

type AuthMiddleware struct {
	tokenManager *token.Manager
}

func NewAuthMiddleware(tokenManager *token.Manager) *AuthMiddleware {
	return &AuthMiddleware{
		tokenManager: tokenManager,
	}
}

func (m *AuthMiddleware) MustLoggedIn() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header is empty",
			})
		}

		if !strings.HasPrefix(authHeader, bearerPrefix) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header format must be Bearer {token}",
			})
		}

		token := strings.TrimSpace(authHeader[len(bearerPrefix):])
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token not found",
			})
		}

		claims, err := m.tokenManager.VerifyAccessToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if claims.Subject == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		c.Locals("userId", claims.Subject)
		return c.Next()
	}
}
