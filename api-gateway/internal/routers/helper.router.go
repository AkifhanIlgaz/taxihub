package routers

import (
	"github.com/AkifhanIlgaz/taxihub/api-gateway/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

type HelperRouter struct {
	helperHandler *handlers.HelperHandler
}

func NewHelperRouter(helperHandler *handlers.HelperHandler) *HelperRouter {
	return &HelperRouter{
		helperHandler: helperHandler,
	}
}

func (r *HelperRouter) RegisterRoutes(api fiber.Router) {
	api.Get("/health", r.helperHandler.Health)
	api.Get("/token", r.helperHandler.CreateToken)
}
