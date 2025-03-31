package routes

import (
	"evermos-app/internal/controllers"
	"evermos-app/internal/middlewares"
	"evermos-app/internal/repositories"

	"github.com/gofiber/fiber/v2"
)

func setupCategoryRoutes(app fiber.Router, ctrl *controllers.CategoryController, jwtSecret string, userRepo repositories.UserRepository) {
	categoryRoute := app.Group("/category")
	categoryRoute.Post("/create", middlewares.AuthMiddleware(jwtSecret, userRepo), middlewares.AdminMiddleware(), ctrl.CreateCategory)
	categoryRoute.Get("/:id", middlewares.AuthMiddleware(jwtSecret, userRepo), middlewares.AdminMiddleware(), ctrl.GetCategory)
	categoryRoute.Get("/", middlewares.AuthMiddleware(jwtSecret, userRepo), middlewares.AdminMiddleware(), ctrl.GetAllCategories)
	categoryRoute.Put("/:id", middlewares.AuthMiddleware(jwtSecret, userRepo), middlewares.AdminMiddleware(), ctrl.UpdateCategory)
	categoryRoute.Delete("/:id", middlewares.AuthMiddleware(jwtSecret, userRepo), middlewares.AdminMiddleware(), ctrl.DeleteCategory)
}
