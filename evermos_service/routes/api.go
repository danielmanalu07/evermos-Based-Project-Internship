package routes

import (
	"evermos-app/config"
	"evermos-app/internal/controllers"
	repositoryimpl "evermos-app/internal/repository_impl"
	"evermos-app/internal/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type API struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewAPI(db *gorm.DB, cfg *config.Config) *API {
	return &API{db: db, cfg: cfg}
}

func (api *API) SetUpRoutes(app *fiber.App) {
	//repositories
	userRepo := repositoryimpl.NewUserRepository(api.db)
	tokoRepo := repositoryimpl.NewTokoRepository(api.db)
	wilayaRepo := repositoryimpl.NewWilayahRepository(api.db)
	alamatRepo := repositoryimpl.NewAlamatRepository(api.db)
	categoryRepo := repositoryimpl.NewCategoryRepository(api.db)
	produkRepo := repositoryimpl.NewProductRepository(api.db)
	trxRepo := repositoryimpl.NewTransactionRepository(api.db)

	//services
	userService := services.NewUserService(userRepo, tokoRepo, wilayaRepo)
	tokoService := services.NewTokoService(tokoRepo, userRepo)
	alamatService := services.NewAlamatService(alamatRepo, userRepo)
	categoryService := services.NewCategoryService(categoryRepo)
	produkService := services.NewProductService(produkRepo, tokoRepo, categoryRepo)
	trxService := services.NewTransactionService(trxRepo, userRepo, alamatRepo, produkRepo, tokoRepo)

	//controller
	userCtrl := controllers.NewUserController(userService, api.cfg.JWTSecret)
	tokoCtrl := controllers.NewTokoController(tokoService)
	alamatCtrl := controllers.NewAlamatController(alamatService)
	categoryCtrl := controllers.NewCategoryController(categoryService)
	produkCtrl := controllers.NewProductController(produkService)
	trxController := controllers.NewTransactionController(trxService)

	apiGroup := app.Group("/api")

	//setup routes
	setupUserRoutes(apiGroup, userCtrl, api.cfg.JWTSecret, userRepo)
	setupTokoRoutes(apiGroup, tokoCtrl, api.cfg.JWTSecret, userRepo)
	setupAlamatRoutes(apiGroup, alamatCtrl, api.cfg.JWTSecret, userRepo)
	setupCategoryRoutes(apiGroup, categoryCtrl, api.cfg.JWTSecret, userRepo)
	setupProductRoutes(apiGroup, produkCtrl, api.cfg.JWTSecret, userRepo)
	setupTransactionRoutes(apiGroup, trxController, api.cfg.JWTSecret, userRepo)
}
