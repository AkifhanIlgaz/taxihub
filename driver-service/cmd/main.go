package main

import (
	"fmt"
	"log"
	"time"

	"github.com/AkifhanIlgaz/taxihub/driver-service/internal/config"
	"github.com/AkifhanIlgaz/taxihub/driver-service/internal/models"
	"github.com/AkifhanIlgaz/taxihub/driver-service/internal/repositories"
	"github.com/AkifhanIlgaz/taxihub/driver-service/internal/services"
	"github.com/AkifhanIlgaz/taxihub/driver-service/pkg/database"
	"github.com/AkifhanIlgaz/taxihub/driver-service/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	mongodb, err := database.ConnectMongo(config.Mongo)
	if err != nil {
		log.Fatal(err)
	}

	err = database.SeedDrivers(mongodb)
	if err != nil {
		log.Fatal(err)
	}

	driverRepo := repositories.NewDriverRepository(mongodb)

	driverService := services.NewDriverService(driverRepo)

	app := fiber.New(fiber.Config{
		AppName:      "TaxiHub Driver Service",
		ServerHeader: "Driver Service",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	app.Use(recover.New())

	driversRouter := app.Group("/drivers")
	driversRouter.Get("/", func(c *fiber.Ctx) error {
		var req models.ListDriversRequest
		if err := c.QueryParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Geçersiz JSON formatı",
			})
		}

		if errors := validator.ValidateStruct(req); errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Doğrulama hatası",
				"errors":  errors,
			})
		}

		drivers, err := driverService.GetDrivers(req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error",
				"error":   err.Error(),
			})
		}
		return c.Status(fiber.StatusFound).JSON(fiber.Map{
			"message": "Sürücüler başarıyla listelendi",
			"data":    drivers,
		})

	})

	driversRouter.Get("/nearby", func(c *fiber.Ctx) error {
		var req models.ListNearbyDriversRequest
		if err := c.QueryParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Geçersiz JSON formatı",
			})
		}

		if errors := validator.ValidateStruct(req); errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Doğrulama hatası",
				"errors":  errors,
			})
		}

		drivers, err := driverService.GetNearbyDrivers(req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error",
				"error":   err.Error(),
			})
		}
		return c.Status(fiber.StatusFound).JSON(fiber.Map{
			"message": "Sürücüler başarıyla listelendi",
			"data":    drivers,
		})

	})

	driversRouter.Post("/", func(c *fiber.Ctx) error {
		var req models.AddDriverRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Geçersiz JSON formatı",
			})
		}

		if errors := validator.ValidateStruct(req); errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Doğrulama hatası",
				"errors":  errors,
			})
		}

		driverId, err := driverService.AddDriver(req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error",
				"error":   err.Error(),
			})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Sürücü başarıyla eklendi",
			"data": fiber.Map{
				"id": driverId,
			},
		})

	})

	driversRouter.Patch("/:id", func(c *fiber.Ctx) error {
		var req models.UpdateDriverRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Geçersiz JSON formatı",
			})
		}
		if errors := validator.ValidateStruct(req); errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Doğrulama hatası",
				"errors":  errors,
			})
		}
		id := c.Params("id")
		if id == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "driver id is empty",
			})
		}

		err = driverService.UpdateDriver(id, req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error",
				"error":   err.Error(),
			})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Sürücü başarıyla guncellendi",
		})
	})

	addr := fmt.Sprintf(":%v", config.Port)
	app.Listen(addr)
}
