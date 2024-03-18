package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tohanilhan/auto-message-sender-service/app/controllers"
	"github.com/tohanilhan/auto-message-sender-service/pkg/middleware"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for GET method:
	route.Get("/messages", middleware.KeyAuthProtected(), controllers.GetMessages) // get list of all messages

	// Routes for POST method:
	route.Post("/change-auto-sending/:option", middleware.KeyAuthProtected(), controllers.ChangeAutoSendingBehaviour) // toggle auto sending messages
}
