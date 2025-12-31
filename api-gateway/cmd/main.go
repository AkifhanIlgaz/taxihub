package main

import (
	"fmt"
	"log"
	"time"

	"github.com/AkifhanIlgaz/taxihub/api-gateway/internal/config"
	"github.com/AkifhanIlgaz/taxihub/api-gateway/internal/middlewares"
	"github.com/AkifhanIlgaz/taxihub/api-gateway/pkg/token"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/gofiber/fiber/v2/middleware/recover"
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

	rateLimiter := middlewares.NewRateLimiter(10, 10*time.Minute)
	logger := middlewares.NewLogger()
	authMiddleware := middlewares.NewAuthMiddleware(tokenManager)

	// accessToken, err := tokenManager.GenerateAccessToken("user_id")
	// if err != nil {
	// 	panic(err)
	// }

	// println(accessToken)

	app := fiber.New(fiber.Config{
		AppName:      "TaxiHub API Gateway",
		ServerHeader: "TaxiHub",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	app.Use(recover.New(), rateLimiter, logger)

	app.Get("/token", func(c *fiber.Ctx) error {
		accessToken, err := tokenManager.GenerateAccessToken("user_id")
		if err != nil {
			panic(err)
		}
		return c.SendString(accessToken)
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "healthy",
			"service": "API Gateway",
			"time":    time.Now().Format(time.RFC3339),
		})
	})

	app.All("/drivers/*", authMiddleware.MustLoggedIn(), func(c *fiber.Ctx) error {
		url := config.DriverService.Url + c.Path()

		if len(c.Request().URI().QueryString()) > 0 {
			url += "?" + string(c.Request().URI().QueryString())
		}

		return proxy.Do(c, url)
	})

	app.Get("/protected", authMiddleware.MustLoggedIn(), func(c *fiber.Ctx) error {
		userId := c.Locals("userId").(string)

		return c.SendString("Hello, " + userId)
	})

	addr := fmt.Sprintf(":%v", config.Port)
	app.Listen(addr)
}
