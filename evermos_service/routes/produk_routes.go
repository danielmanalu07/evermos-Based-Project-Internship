package routes

import (
	"evermos-app/internal/controllers"
	"evermos-app/internal/middlewares"
	"evermos-app/internal/repositories"

	"github.com/gofiber/fiber/v2"
)

func setupProductRoutes(app fiber.Router, ctrl *controllers.ProductController, jwtSecret string, userRepo repositories.UserRepository) {
	app.Get("/page-filter", middlewares.AuthMiddleware(jwtSecret, userRepo), ctrl.GetPageAndFilterProducts)
	poductRoute := app.Group("/product")
	poductRoute.Get("/", middlewares.AuthMiddleware(jwtSecret, userRepo), ctrl.GetAllProducts)
	poductRoute.Post("/create", middlewares.AuthMiddleware(jwtSecret, userRepo), ctrl.CreateProduct)
	poductRoute.Get("/:tokoId", middlewares.AuthMiddleware(jwtSecret, userRepo), ctrl.GetProductsByToko)
	poductRoute.Get("/me/:id", middlewares.AuthMiddleware(jwtSecret, userRepo), ctrl.GetProduct)
	poductRoute.Put("/me/:id", middlewares.AuthMiddleware(jwtSecret, userRepo), ctrl.UpdateProduct)
	poductRoute.Delete("/me/:id", middlewares.AuthMiddleware(jwtSecret, userRepo), ctrl.DeleteProduct)
}
