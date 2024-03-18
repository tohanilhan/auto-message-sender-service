package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/tohanilhan/auto-message-sender-service/docs" // load API Docs files (Swagger)
	"github.com/tohanilhan/auto-message-sender-service/pkg/configs"
	"github.com/tohanilhan/auto-message-sender-service/pkg/middleware"
	"github.com/tohanilhan/auto-message-sender-service/pkg/routes"
	"github.com/tohanilhan/auto-message-sender-service/pkg/utils"
)

func init() {
	// Load .env file.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Parse environment variables.
	err = utils.ParseEnvironmentVariables()
	if err != nil {
		log.Fatal("Error parsing environment variables", err)
	}

	fmt.Printf("%+v\n", utils.EnvironmentVars)
}

// @title						Auto Message Sender Service API
// @version					    1.0
// @description					This is an API for automatic message sender service. Routes are protected with an API key so please use an API key to access private routes. You can find the API key in the .env file of the project. You can also find the API key in the source code of the project on GitHub. The API key should be passed in the header of the request with the key name x-ins-api-auth-key.
// @contact.name				Tohan İlhan
// @contact.email				atahantohan@gmail.com
// @contact.url				https://github.com/tohanilhan/auto-message-sender-service

// @BasePath					/api/v1

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						x-ins-api-auth-key
func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.SwaggerRoute(app)
	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	routes.NotFoundRoute(app)

	// Start server (with or without graceful shutdown).
	if utils.EnvironmentVars.StageStatus == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}

}
