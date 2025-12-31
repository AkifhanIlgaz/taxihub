package routers

import (
	"github.com/AkifhanIlgaz/taxihub/api-gateway/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

type DriverRouter struct {
	driverHandler *handlers.DriverHandler
}

func NewDriverRouter(driverHandler *handlers.DriverHandler) *DriverRouter {
	return &DriverRouter{
		driverHandler: driverHandler,
	}
}

func (r *DriverRouter) RegisterRoutes(api fiber.Router) {
	drivers := api.Group("/drivers")

	drivers.Get("/", r.driverHandler.GetDrivers)
	drivers.Get("/nearby", r.driverHandler.GetNearbyDrivers)
	drivers.Post("/", r.driverHandler.AddDriver)
	drivers.Patch("/:id", r.driverHandler.UpdateDriver)
}
