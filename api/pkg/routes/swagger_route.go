package routes

import (
	"github.com/gofiber/fiber/v2"

	swagger "github.com/gofiber/swagger"
)

// SwaggerRoute func for describe group of API Docs routes.
func SwaggerRoute(a *fiber.App) {
	// Create routes group.
	route := a.Group("api/v1")

	// Routes for GET method:
	route.Get("/swagger/*", swagger.HandlerDefault)
}
