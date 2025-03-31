package routes

import (
	"evermos-app/internal/controllers"
	"evermos-app/internal/middlewares"
	"evermos-app/internal/repositories"

	"github.com/gofiber/fiber/v2"
)

func setupTokoRoutes(app fiber.Router, ctrl *controllers.TokoController, jwtSecret string, userRepo repositories.UserRepository) {
	tokoRoute := app.Group("/toko")
	tokoRoute.Get("/me", middlewares.AuthMiddleware(jwtSecret, userRepo), ctrl.GetUserTokos)
	tokoRoute.Get("/me/:id", middlewares.AuthMiddleware(jwtSecret, userRepo), ctrl.GetToko)
	tokoRoute.Put("/me/:id", middlewares.AuthMiddleware(jwtSecret, userRepo), ctrl.UpdateToko)
	tokoRoute.Delete("/me/:id", middlewares.AuthMiddleware(jwtSecret, userRepo), ctrl.DeleteToko)
}
