package main

import (
	"fmt"
	"log"
	"time"

	"github.com/AkifhanIlgaz/taxihub/api-gateway/internal/clients"
	"github.com/AkifhanIlgaz/taxihub/api-gateway/internal/config"
	"github.com/AkifhanIlgaz/taxihub/api-gateway/internal/handlers"
	"github.com/AkifhanIlgaz/taxihub/api-gateway/internal/middlewares"
	"github.com/AkifhanIlgaz/taxihub/api-gateway/internal/routers"
	"github.com/AkifhanIlgaz/taxihub/api-gateway/pkg/token"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	tokenManager, err := token.NewManager(config.Token)
	if err != nil {
		log.Fatal(err)
	}

	clientManager, err := clients.NewClientManager(config.DriverService.Url)
	if err != nil {
		log.Fatal(err)
	}

	authMiddleware := middlewares.NewAuthMiddleware(tokenManager)
	rateLimiter := middlewares.NewRateLimiter(10, 10*time.Minute)
	logger := middlewares.NewLogger()

	driverHandler := handlers.NewDriverHandler(clientManager.DriverClient)

	driverRouter := routers.NewDriverRouter(driverHandler)

	app := fiber.New(fiber.Config{
		AppName:      "TaxiHub API Gateway",
		ServerHeader: "TaxiHub",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	api := app.Group("/api", authMiddleware.MustLoggedIn(), rateLimiter, logger)

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "healthy",
			"service": "API Gateway",
			"time":    time.Now().Format(time.RFC3339),
		})
	})

	// JWT token test etmek icin
	api.Get("/token", func(c *fiber.Ctx) error {
		accessToken, err := tokenManager.GenerateAccessToken("user_123")
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(fiber.Map{"token": accessToken})
	})

	driverRouter.RegisterRoutes(api)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"error":   "Route not found",
		})
	})

	addr := fmt.Sprintf(":%v", config.Port)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
