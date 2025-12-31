package handlers

import (
	"context"
	"errors"
	"time"

	dto "github.com/AkifhanIlgaz/taxihub/api-gateway/internal/dto/request"
	"github.com/AkifhanIlgaz/taxihub/api-gateway/pkg/response"
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
		return response.Error(c, fiber.StatusBadRequest, errors.New("Invalid query parameters"))
	}

	if errs := validator.ValidateStruct(req); errs != nil {
		if len(errs) == 0 {
			return response.Error(c, fiber.StatusBadRequest, errors.New("something went wrong with validation"))
		}
		return response.Error(c, fiber.StatusBadRequest, errs[0])
	}

	protoReq := dto.GetDriversRequestToProto(&req)
	protoResp, err := h.client.GetDrivers(ctx, protoReq)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, protoResp, "Sürücüler başarıyla listelendi")
}

func (h *DriverHandler) GetNearbyDrivers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var req dto.GetNearbyDriversRequest
	if err := c.QueryParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, errors.New("Invalid query parameters"))
	}

	if errs := validator.ValidateStruct(req); errs != nil {
		if len(errs) == 0 {
			return response.Error(c, fiber.StatusBadRequest, errors.New("something went wrong with validation"))
		}
		return response.Error(c, fiber.StatusBadRequest, errs[0])
	}

	protoReq := dto.GetNearbyDriversRequestToProto(&req)
	protoResp, err := h.client.GetNearbyDrivers(ctx, protoReq)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, protoResp, "Sürücüler başarıyla listelendi")
}

func (h *DriverHandler) AddDriver(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var req dto.AddDriverRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, errors.New("Invalid body parameters"))
	}

	if errs := validator.ValidateStruct(req); errs != nil {
		if len(errs) == 0 {
			return response.Error(c, fiber.StatusBadRequest, errors.New("something went wrong with validation"))
		}
		return response.Error(c, fiber.StatusBadRequest, errs[0])
	}

	protoReq := dto.AddDriverRequestToProto(&req)
	protoResp, err := h.client.AddDriver(ctx, protoReq)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, protoResp, "Sürücü başarıyla eklendi")
}

func (h *DriverHandler) UpdateDriver(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var req dto.UpdateDriverRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, errors.New("Invalid body parameters"))
	}

	if errs := validator.ValidateStruct(req); errs != nil {
		if len(errs) == 0 {
			return response.Error(c, fiber.StatusBadRequest, errors.New("something went wrong with validation"))
		}
		return response.Error(c, fiber.StatusBadRequest, errs[0])
	}

	id := c.Params("id")
	if id == "" {
		return response.Error(c, fiber.StatusBadRequest, errors.New("driver id is empty"))
	}

	protoReq := dto.UpdateDriverRequestToProto(&req)
	protoReq.Id = id

	protoResp, err := h.client.UpdateDriver(ctx, protoReq)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, protoResp, "Sürücü başarıyla güncellendi")
}
