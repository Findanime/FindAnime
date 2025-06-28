package routes

import (
	"api/internal/middleware"
	routes "api/internal/routes/public"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	app.Use(middleware.Logging)

	// Public API Routes
	publicGroup := app.Group("/api/v1/public")
	publicGroup.Get("/recommend", routes.Recommend)
}
