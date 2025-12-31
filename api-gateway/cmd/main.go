package main

import (
	"fmt"
	"log"
	"time"

	"github.com/AkifhanIlgaz/taxihub/api-gateway/internal/clients"
	"github.com/AkifhanIlgaz/taxihub/api-gateway/internal/config"
	"github.com/AkifhanIlgaz/taxihub/api-gateway/internal/routes"
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

	app := fiber.New(fiber.Config{
		AppName:      "TaxiHub API Gateway",
		ServerHeader: "TaxiHub",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	routes.Setup(app, clientManager, tokenManager, config)

	addr := fmt.Sprintf(":%v", config.Port)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
