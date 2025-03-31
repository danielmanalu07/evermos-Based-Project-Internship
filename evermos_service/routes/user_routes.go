package routes

import (
	"evermos-app/internal/controllers"
	"evermos-app/internal/middlewares"
	"evermos-app/internal/repositories"

	"github.com/gofiber/fiber/v2"
)

func setupUserRoutes(app fiber.Router, ctrl *controllers.UserController, jwtSecret string, userRepo repositories.UserRepository) {
	app.Post("/register", ctrl.RegisterUser)
	app.Post("/login", ctrl.LoginUser)

	userRoute := app.Group("/users")
	userRoute.Use(middlewares.AuthMiddleware(jwtSecret, userRepo))
	userRoute.Get("/me", ctrl.GetProfile)
	userRoute.Put("/me/update", ctrl.UpdateProfile)
	userRoute.Post("/me/logout", ctrl.LogoutUser)
	userRoute.Delete("/me/delete", ctrl.DeleteAccount)
}
