package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestPublicRoutes(t *testing.T) {

	if err := godotenv.Load("../../test.env"); err != nil {
		panic(err)
	}

	tests := []struct {
		description   string
		route         string // input route
		expectedError bool
		expectedCode  int
	}{
		{
			description:   "Ping route",
			route:         "/api/v1/ping",
			expectedError: false,
			expectedCode:  200,
		},
	}

	// Define Fiber app.
	app := fiber.New()

	// Define routes.
	PublicRoutes(app)

	for _, test := range tests {
		// Create a new http request with the route from the test case.
		req := httptest.NewRequest("GET", test.route, http.NoBody)
		req.Header.Set("Content-Type", "application/json")

		// Perform the request plain with the app.
		resp, err := app.Test(req, -1) // the -1 disables request latency

		// Verify, that no error occurred, that is not expected
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		// As expected errors lead to broken responses,
		// the next test case needs to be processed.
		if test.expectedError {
			continue
		}

		// Verify, if the status code is as expected.
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
