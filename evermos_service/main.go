package main

import (
	"evermos-app/config"
	"evermos-app/pkg/databases"
	"evermos-app/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.LoadConfig()

	db, err := databases.ConnectDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := databases.Migrate(db); err != nil {
		log.Fatal("Failed to migrate to database:", err)
	}

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "https://gofiber.io",
	}))

	api := routes.NewAPI(db, cfg)
	api.SetUpRoutes(app)

	log.Fatal(app.Listen(":" + cfg.Port))
}
