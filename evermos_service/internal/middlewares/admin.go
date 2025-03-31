package middlewares

import (
	"evermos-app/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func AdminMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		isAdmin := c.Locals("isAdmin").(bool)
		if !isAdmin {
			return utils.ErrorResponse(c, fiber.StatusForbidden, "Admin Access required")
		}

		return c.Next()
	}
}
