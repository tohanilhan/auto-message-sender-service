package utils

import (
	"testing"
)

func TestConnectionURLBuilder(t *testing.T) {

	// Prepare test environment variables
	mockEnvironmentVariables()

	// Clean up environment variables after the test
	defer unSetEnvVars()

	// Call the function being tested
	err := ParseEnvironmentVariables()

	// Check if there's any error
	if err != nil {
		t.Fatalf("ParseEnvironmentVariables failed: %v", err)
	}

	// Test cases
	testCases := []struct {
		name        string
		connection  string
		expectedURL string
		expectError bool
	}{
		{
			name:        "PostgresConnection",
			connection:  "postgres",
			expectedURL: "host=localhost port=5432 user=user password=password dbname=testdb sslmode=disable",
			expectError: false,
		},
		{
			name:        "RedisConnection",
			connection:  "redis",
			expectedURL: "localhost:6379",
			expectError: false,
		},
		{
			name:        "InvalidConnection",
			connection:  "invalid",
			expectedURL: "",
			expectError: true,
		},
		{
			name:        "FiberConnection",
			connection:  "fiber",
			expectedURL: ":8080",
			expectError: false,
		},
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the function being tested
			url, err := ConnectionURLBuilder(tc.connection)

			// Check if there's any error
			if (err != nil) != tc.expectError {
				t.Fatalf("Test case '%s': unexpected error state. Expected error: %t, Got error: %t", tc.name, tc.expectError, err != nil)
			}

			// Check if the built URL matches the expected URL
			if url != tc.expectedURL {
				t.Errorf("Test case '%s': unexpected URL. Expected: %s, Got: %s", tc.name, tc.expectedURL, url)
			}
		})
	}
}
