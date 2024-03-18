package configs

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tohanilhan/auto-message-sender-service/pkg/utils"
)

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func FiberConfig() fiber.Config {
	// Define server settings.

	// Return Fiber configuration.
	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(utils.EnvironmentVars.ServerReadTimeout),
	}
}
