package main

import (
	"api/internal/routes"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	// Setup Routes
	routes.SetupRoutes(app)
	app.Listen(":3001")
}
