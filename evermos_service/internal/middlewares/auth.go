package middlewares

import (
	"evermos-app/internal/repositories"
	"evermos-app/internal/utils"
	"evermos-app/pkg/auth"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(jwtSecret string, userRepo repositories.UserRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized")
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized: Invalid token format. Use 'Bearer <token>'")
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		if userRepo.IsTokenRevoked(token) {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Token has been revoked")
		}

		claims, err := auth.ParseToken(token, jwtSecret)
		if err != nil {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid Token")
		}

		c.Locals("userId", claims.UserID)
		c.Locals("isAdmin", claims.IsAdmin)
		return c.Next()
	}
}
