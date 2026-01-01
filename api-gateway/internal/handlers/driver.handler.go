package handlers

import (
	"context"
	"errors"
	"time"

	dtoReq "github.com/AkifhanIlgaz/taxihub/api-gateway/internal/dto/request"
	"github.com/AkifhanIlgaz/taxihub/api-gateway/internal/mapper"

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

// @Summary Driver'ları listele
// @Tags driver
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page"
// @Param pageSize query int false "Page Size"
// @Success 200 {object} response.APIResponse{data=dto.GetDriversResponse}
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /drivers [get]
func (h *DriverHandler) GetDrivers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var req dtoReq.GetDriversRequest
	if err := c.QueryParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, errors.New("Invalid query parameters"))
	}

	if errs := validator.ValidateStruct(req); errs != nil {
		if len(errs) == 0 {
			return response.Error(c, fiber.StatusBadRequest, errors.New("something went wrong with validation"))
		}
		return response.Error(c, fiber.StatusBadRequest, errs[0])
	}

	protoReq := req.ToProto()
	protoResp, err := h.client.GetDrivers(ctx, protoReq)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	res := mapper.ProtoToGetDriversResponse(protoResp)

	return response.Success(c, res, "Sürücüler başarıyla listelendi")
}

// @Summary Verilen koordinatların 6km yakınındaki driverları listeler
// @Tags driver
// @Produce json
// @Security BearerAuth
// @Param lat query number true "Latitude"
// @Param lon query number true "Longitude"
// @Param taxiType query string true "Taxi Type" Enums(sari, korsan, uber, tag)
// @Success 200 {object} response.APIResponse{data=dto.GetNearbyDriversResponse}
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /drivers/nearby [get]
func (h *DriverHandler) GetNearbyDrivers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var req dtoReq.GetNearbyDriversRequest
	if err := c.QueryParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, errors.New("Invalid query parameters"))
	}

	if errs := validator.ValidateStruct(req); errs != nil {
		if len(errs) == 0 {
			return response.Error(c, fiber.StatusBadRequest, errors.New("something went wrong with validation"))
		}
		return response.Error(c, fiber.StatusBadRequest, errs[0])
	}

	protoReq := req.ToProto()
	protoResp, err := h.client.GetNearbyDrivers(ctx, protoReq)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	res := mapper.ProtoToGetNearbyDriversResponse(protoResp)

	return response.Success(c, res, "Sürücüler başarıyla listelendi")
}

// @Summary Driver ekler
// @Tags driver
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body dto.AddDriverRequest true "Add driver"
// / @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /drivers [post]
func (h *DriverHandler) AddDriver(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var req dtoReq.AddDriverRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, errors.New("Invalid body parameters"))
	}

	if errs := validator.ValidateStruct(req); errs != nil {
		if len(errs) == 0 {
			return response.Error(c, fiber.StatusBadRequest, errors.New("something went wrong with validation"))
		}
		return response.Error(c, fiber.StatusBadRequest, errs[0])
	}

	protoReq := req.ToProto()
	protoResp, err := h.client.AddDriver(ctx, protoReq)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, protoResp, "Driver başarıyla eklendi")
}

// @Summary Driver günceller
// @Tags driver
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Driver ID"
// @Param body body dto.UpdateDriverRequest true "Update driver"
// @Success 200 {object} response.APIResponse
// @Failure 400 {object} response.APIResponse
// @Failure 500 {object} response.APIResponse
// @Router /drivers/{id} [put]
func (h *DriverHandler) UpdateDriver(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var req dtoReq.UpdateDriverRequest
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

	protoReq := req.ToProto()
	protoReq.Id = id

	protoResp, err := h.client.UpdateDriver(ctx, protoReq)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err)
	}

	return response.Success(c, protoResp, "Driver başarıyla güncellendi")
}
