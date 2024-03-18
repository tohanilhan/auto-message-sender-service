package routes

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/tohanilhan/auto-message-sender-service/pkg/utils"
)

func TestPrivateRoutes(t *testing.T) {

	if err := godotenv.Load("../../test.env"); err != nil {
		panic(err)
	}
	// Parse environment variables.
	err := utils.ParseEnvironmentVariables()
	if err != nil {
		log.Fatal("Error parsing environment variables", err)
	}

	tests := []struct {
		description           string
		method                string // input method
		route                 string // input route
		param                 string // input params
		apiHeaderKey          string
		apiHeaderValue        string
		expectedError         bool
		expectedResponseError bool
		expectedCode          int
	}{
		{
			description:           "Get all messages route",
			method:                "GET",
			route:                 "/api/v1/messages",
			apiHeaderKey:          "x-ins-api-auth-key",
			apiHeaderValue:        "INS.er7u2oVtHsmlqWICxMnF.pD5k8zLcYh3iR6XaO",
			expectedError:         false,
			expectedResponseError: false,
			expectedCode:          200,
		},
		{
			description:           "Change auto sending behaviour route",
			method:                "POST",
			route:                 "/api/v1/change-auto-sending/",
			param:                 "ON",
			apiHeaderKey:          "x-ins-api-auth-key",
			apiHeaderValue:        "INS.er7u2oVtHsmlqWICxMnF.pD5k8zLcYh3iR6XaO",
			expectedError:         false,
			expectedResponseError: false,
			expectedCode:          200,
		},
		{
			description:           "Change auto sending behaviour route",
			method:                "POST",
			route:                 "/api/v1/change-auto-sending/",
			param:                 "OFF",
			apiHeaderKey:          "x-ins-api-auth-key",
			apiHeaderValue:        "INS.er7u2oVtHsmlqWICxMnF.pD5k8zLcYh3iR6XaO",
			expectedError:         false,
			expectedResponseError: false,
			expectedCode:          200,
		},
		{
			description:           "Change auto sending behaviour route",
			method:                "POST",
			route:                 "/api/v1/change-auto-sending/",
			param:                 "INVALID",
			apiHeaderKey:          "x-ins-api-auth-key",
			apiHeaderValue:        "INS.er7u2oVtHsmlqWICxMnF.pD5k8zLcYh3iR6XaO",
			expectedError:         false,
			expectedResponseError: true,
			expectedCode:          400,
		},
		{
			description:           "Change auto sending behaviour route",
			method:                "POST",
			route:                 "/api/v1/change-auto-sending/",
			param:                 "ON",
			apiHeaderKey:          "x-ins-api-auth-key",
			apiHeaderValue:        "INVALID",
			expectedError:         false,
			expectedResponseError: true,
			expectedCode:          401,
		},
		{
			description:           "Change auto sending behaviour route",
			method:                "POST",
			route:                 "/api/v1/change-auto-sending/",
			param:                 "ON",
			apiHeaderKey:          "INVALID",
			apiHeaderValue:        "INS.er7u2oVtHsmlqWICxMnF.pD5k8zLcYh3iR6XaO",
			expectedError:         false,
			expectedResponseError: false,
			expectedCode:          401,
		},
		{
			description:    "Change auto sending behaviour route",
			method:         "GET",
			route:          "/api/v1/change-auto-sending/",
			apiHeaderKey:   "x-ins-api-auth-key",
			apiHeaderValue: "INS.er7u2oVtHsmlqWICxMnF.pD5k8zLcYh3iR6XaO",
			expectedError:  false,
			expectedCode:   404,
		},
		{
			description:           "Get all messages route",
			method:                "GET",
			route:                 "/api/v1/messages",
			apiHeaderKey:          "x-ins-api-auth-key",
			apiHeaderValue:        "INVALID",
			expectedError:         false,
			expectedResponseError: true,
			expectedCode:          401,
		},
		{
			description:           "Get all messages route",
			method:                "GET",
			route:                 "/api/v1/messages",
			apiHeaderKey:          "x-ins-api-auth-key",
			apiHeaderValue:        "INVALID",
			expectedError:         false,
			expectedResponseError: true,
			expectedCode:          401,
		},
		{
			description:           "Get all messages route",
			method:                "GET",
			route:                 "/api/v1/messages",
			apiHeaderKey:          "INVALID",
			apiHeaderValue:        "INS.er7u2oVtHsmlqWICxMnF.pD5k8zLcYh3iR6XaO",
			expectedError:         false,
			expectedResponseError: true,
			expectedCode:          401,
		},
		{
			description:           "Get all messages route",
			method:                "POST",
			route:                 "/api/v1/messages",
			apiHeaderKey:          "x-ins-api-auth-key",
			apiHeaderValue:        "INS.er7u2oVtHsmlqWICxMnF.pD5k8zLcYh3iR6XaO",
			expectedError:         false,
			expectedResponseError: true,
			expectedCode:          405,
		},
	}
	// Define Fiber app.
	app := fiber.New()

	// Define routes.
	PrivateRoutes(app)

	for _, test := range tests {
		// Create a new http request with the route from the test case.
		if test.param != "" {
			test.route += test.param
		}

		req := httptest.NewRequest(test.method, test.route, http.NoBody)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set(test.apiHeaderKey, test.apiHeaderValue)

		// Perform the request plain with the app.
		resp, err := app.Test(req, -1) // the -1 disables request latency

		// Verify, that no error occurred, that is not expected
		assert.Equalf(t, test.expectedError, err != nil, test.description)

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
