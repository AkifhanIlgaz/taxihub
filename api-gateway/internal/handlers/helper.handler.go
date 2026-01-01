package handlers

import (
	"time"

	dto "github.com/AkifhanIlgaz/taxihub/api-gateway/internal/dto/response"
	"github.com/AkifhanIlgaz/taxihub/api-gateway/pkg/token"
	"github.com/gofiber/fiber/v2"
)

type HelperHandler struct {
	tokenManager *token.Manager
}

func NewHelperHandler(tokenManager *token.Manager) *HelperHandler {
	return &HelperHandler{tokenManager: tokenManager}
}

// Health godoc
// @Summary Health check
// @Description Returns API Gateway health status with timestamp
// @Tags helper
// @Produce json
// @Success 200 {object} dto.HealthResponse
// @Router /health [get]
func (h *HelperHandler) Health(c *fiber.Ctx) error {
	return c.JSON(dto.HealthResponse{
		Status:  "healthy",
		Service: "API Gateway",
		Time:    time.Now().Format(time.RFC3339),
	})
}

// CreateToken godoc
// @Summary Create access token
// @Description Generates a mock access token for testing
// @Tags helper
// @Produce json
// @Success 200 {object} dto.CreateTokenResponse
// @Failure 500 {object} map[string]string
// @Router /token [get]
func (h *HelperHandler) CreateToken(c *fiber.Ctx) error {
	accessToken, err := h.tokenManager.GenerateAccessToken("mock_user_id")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(dto.CreateTokenResponse{Token: accessToken})
}
