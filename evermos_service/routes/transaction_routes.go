package routes

import (
	"evermos-app/internal/controllers"
	"evermos-app/internal/middlewares"
	"evermos-app/internal/repositories"

	"github.com/gofiber/fiber/v2"
)

func setupTransactionRoutes(app fiber.Router, ctrl *controllers.TransactionController, jwtSecret string, userRepo repositories.UserRepository) {
	app.Get("/page-filter", middlewares.AuthMiddleware(jwtSecret, userRepo), ctrl.GetPageAndFilterTransaksis)
	trxRoute := app.Group("/transaksi")
	trxRoute.Post("/create", middlewares.AuthMiddleware(jwtSecret, userRepo), ctrl.CreateTransaksi)
	trxRoute.Get("/:id", middlewares.AuthMiddleware(jwtSecret, userRepo), ctrl.GetTransaksi)
	trxRoute.Get("/", middlewares.AuthMiddleware(jwtSecret, userRepo), ctrl.GetAllTransaksis)
}
