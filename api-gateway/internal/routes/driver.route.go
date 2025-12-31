package routes

import (
	"time"

	"github.com/AkifhanIlgaz/taxihub/api-gateway/internal/clients"
	"github.com/AkifhanIlgaz/taxihub/api-gateway/internal/config"
	"github.com/AkifhanIlgaz/taxihub/api-gateway/internal/handlers"
	"github.com/AkifhanIlgaz/taxihub/api-gateway/internal/middlewares"
	"github.com/AkifhanIlgaz/taxihub/api-gateway/pkg/token"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, clientManager *clients.ClientManager, tokenManager *token.Manager, cfg config.ServiceConfig) {
	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "healthy",
			"service": "API Gateway",
			"time":    time.Now().Format(time.RFC3339),
		})
	})

	// Test token endpoint (remove in production)
	app.Get("/token", func(c *fiber.Ctx) error {
		accessToken, err := tokenManager.GenerateAccessToken("user_123")
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(fiber.Map{"token": accessToken})
	})

	api := app.Group("/api")

	// Protected routes
	authMiddleware := middlewares.NewAuthMiddleware(tokenManager)
	rateLimiter := middlewares.NewRateLimiter(10, 10*time.Minute)
	logger := middlewares.NewLogger()

	protected := api.Group("", authMiddleware.MustLoggedIn(), rateLimiter, logger)

	// Driver routes
	driverHandler := handlers.NewDriverHandler(clientManager.DriverClient)
	drivers := protected.Group("/drivers")
	drivers.Get("/", driverHandler.GetDrivers)
	drivers.Get("/nearby", driverHandler.GetNearbyDrivers)
	drivers.Post("/", driverHandler.AddDriver)
	drivers.Put("/:id", driverHandler.UpdateDriver)

	// 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   "Route not found",
		})
	})
}
