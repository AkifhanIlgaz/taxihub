package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewLogger() fiber.Handler {
	return logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${method} ${path} (${latency}) - IP: ${ip} - User-Agent: ${ua} - User ID: ${userId}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
		Output:     os.Stdout, //  bir dosyaya veya elasticsearch e yazilabilir
		CustomTags: map[string]logger.LogFunc{
			"userId": func(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
				if userId := c.Locals("userId"); userId != nil {
					return output.WriteString(userId.(string))
				}
				return output.WriteString("Anonymous")
			},
		},
	})
}
