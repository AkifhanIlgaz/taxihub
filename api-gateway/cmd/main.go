package main

import (
	"fmt"
	"log"
	"time"

	_ "github.com/AkifhanIlgaz/taxihub/api-gateway/docs"
	"github.com/AkifhanIlgaz/taxihub/api-gateway/internal/clients"
	"github.com/AkifhanIlgaz/taxihub/api-gateway/internal/config"
	"github.com/AkifhanIlgaz/taxihub/api-gateway/internal/handlers"
	"github.com/AkifhanIlgaz/taxihub/api-gateway/internal/middlewares"
	"github.com/AkifhanIlgaz/taxihub/api-gateway/internal/routers"
	"github.com/AkifhanIlgaz/taxihub/api-gateway/pkg/token"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title TaxiHub API Gateway
// @version 1.0
// @description TaxiHub API Gateway dokumantasyonu
// @host localhost:8080
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Authorization header must start with "Bearer ".
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
	rateLimiter := middlewares.NewRateLimiter(config.RateLimit)
	logger := middlewares.NewLogger()

	driverHandler := handlers.NewDriverHandler(clientManager.DriverClient)
	helperHandler := handlers.NewHelperHandler(tokenManager)

	driverRouter := routers.NewDriverRouter(driverHandler)
	helperRouter := routers.NewHelperRouter(helperHandler)

	app := fiber.New(fiber.Config{
		AppName:      "TaxiHub API Gateway",
		ServerHeader: "TaxiHub",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	app.Get("/swagger/*", swagger.HandlerDefault)

	api := app.Group("/api", rateLimiter, logger)
	helperRouter.RegisterRoutes(api)

	protectedApi := api.Group("", authMiddleware.MustLoggedIn())

	driverRouter.RegisterRoutes(protectedApi)

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
