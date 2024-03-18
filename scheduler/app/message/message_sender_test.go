package message

import (
	"encoding/json"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/tohanilhan/auto-message-sender-service/scheduler/app/models"
	"github.com/tohanilhan/auto-message-sender-service/scheduler/pkg/utils"
)

// Unit tests for handleRequest function
func TestHandleRequest(t *testing.T) {
	// Load .env file.
	err := godotenv.Load("../../test.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Parse environment variables.
	err = utils.ParseEnvironmentVariables()
	if err != nil {
		log.Fatal("Error parsing environment variables")
	}

	// Mock MessageSender
	mockMessageSender := &MessageSender{WebhookURL: utils.EnvironmentVars.WebhookApiUrl, SentTime: time.Now()}

	// Define a test message
	testMessage := models.Message{Content: "Test Message", To: "+1234567890"}

	// Call the function being tested
	response, err := mockMessageSender.handleRequest(testMessage)

	// Check if there's any error
	if err != nil {
		t.Errorf("handleRequest function failed: %v", err)
	}

	// Validate the response format
	var respData map[string]interface{}
	err = json.Unmarshal(response, &respData)
	if err != nil {
		t.Errorf("Failed to unmarshal response JSON: %v", err)
	}

	// Check if the response contains expected fields
	if _, ok := respData["message"]; !ok {
		t.Error("Response is missing 'status' field")
	}

	if _, ok := respData["messageId"]; !ok {
		t.Error("Response is missing 'message' field")
	}

}

// func TestUnmarshalResponse(t *testing.T) {
// 	// Mock MessageSender
// 	mockMessageSender := &MessageSender{}

// 	// Define a test response
// 	testResponse := `{"message":"Accepted","messageId":"7d3c5c76-d41b-4f3d-95d5-7d3fbef92dc1"}`

// 	// Call the function being tested
// 	responseStruct, err := mockMessageSender.unmarshalResponse(testResponse)

// 	// Check if there's any error
// 	if err != nil {
// 		t.Errorf("unmarshalResponse function failed: %v", err)
// 	}

// 	// Validate the response structure
// 	expectedResponse := &models.SendMessageResponse{
// 		Message:   "Accepted",
// 		MessageID: "7d3c5c76-d41b-4f3d-95d5-7d3fbef92dc1",
// 	}

// 	// Compare the response with the expected structure
// 	if !reflect.DeepEqual(responseStruct, expectedResponse) {
// 		t.Errorf("unmarshalResponse returned unexpected response. Expected: %+v, Got: %+v", expectedResponse, responseStruct)
// 	}
// }

func TestUnmarshalResponse(t *testing.T) {
	// Define test cases with different response structures
	testCases := []struct {
		name           string
		response       string
		expectedResult *models.SendMessageResponse
		expectedError  bool
	}{
		{
			name:     "ValidResponse",
			response: `{"message":"Accepted","messageId":"7d3c5c76-d41b-4f3d-95d5-7d3fbef92dc1"}`,
			expectedResult: &models.SendMessageResponse{
				Message:   "Accepted",
				MessageID: "7d3c5c76-d41b-4f3d-95d5-7d3fbef92dc1",
			},
			expectedError: false,
		},
		{
			name:           "EmptyResponse",
			response:       "",
			expectedResult: nil,
			expectedError:  true,
		},
		{
			name:           "InvalidJSON",
			response:       `{"status":"success","message":"Message sent successfully","id":"123456",}`,
			expectedResult: nil,
			expectedError:  true,
		},
	}

	// Iterate over test cases
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Mock MessageSender
			mockMessageSender := &MessageSender{}

			// Call the function being tested
			responseStruct, err := mockMessageSender.unmarshalResponse(testCase.response)

			// Check if there's any error
			if (err != nil) != testCase.expectedError {
				t.Errorf("Test case '%s': unexpected error state. Expected error: %t, Got error: %t", testCase.name, testCase.expectedError, err != nil)
			}

			// Compare the response with the expected structure
			if !reflect.DeepEqual(responseStruct, testCase.expectedResult) {
				t.Errorf("Test case '%s': unexpected response. Expected: %+v, Got: %+v", testCase.name, testCase.expectedResult, responseStruct)
			}
		})
	}
}
