package main

import (
	"fmt"
	"log"
	"time"

	"github.com/AkifhanIlgaz/taxihub/api-gateway/internal/config"
	"github.com/AkifhanIlgaz/taxihub/api-gateway/internal/middlewares"
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

	rateLimiter := middlewares.NewRateLimiter(2, 10*time.Minute)
	logger := middlewares.NewLogger()
	authMiddleware := middlewares.NewAuthMiddleware(tokenManager)

	accessToken, err := tokenManager.GenerateAccessToken("user_id")
	if err != nil {
		panic(err)
	}

	println(accessToken)

	app := fiber.New()
	app.Use(rateLimiter, logger)

	app.Get("/token", func(c *fiber.Ctx) error {
		return c.SendString("public")
	})

	app.Get("/protected", authMiddleware.MustLoggedIn(), func(c *fiber.Ctx) error {
		userId := c.Locals("userId").(string)

		return c.SendString("Hello, " + userId)
	})

	addr := fmt.Sprintf(":%v", config.Port)
	app.Listen(addr)
}
