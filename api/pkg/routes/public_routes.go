package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tohanilhan/auto-message-sender-service/app/controllers"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for GET method:
	route.Get("/ping", controllers.Ping) // check if the server is alive
}
