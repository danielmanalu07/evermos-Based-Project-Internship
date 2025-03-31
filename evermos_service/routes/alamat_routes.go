package routes

import (
	"evermos-app/internal/controllers"
	"evermos-app/internal/middlewares"
	"evermos-app/internal/repositories"

	"github.com/gofiber/fiber/v2"
)

func setupAlamatRoutes(app fiber.Router, ctrl *controllers.AlamatController, jwtSecret string, userRepo repositories.UserRepository) {
	alamatRoute := app.Group("/alamat")
	alamatRoute.Post("/me/create", middlewares.AuthMiddleware(jwtSecret, userRepo), ctrl.CreateAlamat)
	alamatRoute.Get("/me", middlewares.AuthMiddleware(jwtSecret, userRepo), ctrl.GetUserAlamats)
	alamatRoute.Get("/me/:id", middlewares.AuthMiddleware(jwtSecret, userRepo), ctrl.GetAlamat)
	alamatRoute.Put("/me/:id", middlewares.AuthMiddleware(jwtSecret, userRepo), ctrl.UpdateAlamat)
	alamatRoute.Delete("/me/:id", middlewares.AuthMiddleware(jwtSecret, userRepo), ctrl.DeleteAlamat)
}
