package handlers

import (
	"context"
	"time"

	dto "github.com/AkifhanIlgaz/taxihub/api-gateway/internal/dto/request"
	"github.com/AkifhanIlgaz/taxihub/api-gateway/pkg/validator"
	pb "github.com/AkifhanIlgaz/taxihub/common/proto/driver"
	"github.com/gofiber/fiber/v2"
)

type DriverHandler struct {
	client pb.DriverServiceClient
}

func NewDriverHandler(client pb.DriverServiceClient) *DriverHandler {
	return &DriverHandler{client: client}
}

func (h *DriverHandler) GetDrivers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var req dto.GetDriversRequest
	if err := c.QueryParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid query parameters",
		})
	}

	if errors := validator.ValidateStruct(req); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Doğrulama hatası",
			"errors":  errors,
		})
	}

	protoReq := dto.GetDriversRequestToProto(&req)
	protoResp, err := h.client.GetDrivers(ctx, protoReq)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusFound).JSON(fiber.Map{
		"message": "Sürücüler başarıyla listelendi",
		"data":    protoResp,
	})
}

func (h *DriverHandler) GetNearbyDrivers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var req dto.GetNearbyDriversRequest
	if err := c.QueryParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid query parameters",
		})
	}

	if errors := validator.ValidateStruct(req); errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Doğrulama hatası",
			"errors":  errors,
		})
	}

	protoReq := dto.GetNearbyDriversRequestToProto(&req)
	protoResp, err := h.client.GetNearbyDrivers(ctx, protoReq)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusFound).JSON(fiber.Map{
		"message": "Sürücüler başarıyla listelendi",
		"data":    protoResp,
	})
}

func (h *DriverHandler) AddDriver(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var req dto.AddDriverRequest
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

	protoReq := dto.AddDriverRequestToProto(&req)
	protoResp, err := h.client.AddDriver(ctx, protoReq)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Sürücü başarıyla eklendi",
		"data":    protoResp,
	})
}

func (h *DriverHandler) UpdateDriver(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var req dto.UpdateDriverRequest
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

	protoReq := dto.UpdateDriverRequestToProto(&req)
	protoResp, err := h.client.UpdateDriver(ctx, protoReq)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": protoResp,
	})
}
