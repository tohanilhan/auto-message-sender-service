package utils

import (
	"os"
	"testing"
)

func mockEnvironmentVariables() {
	// Prepare test environment variables
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "user")
	os.Setenv("DB_PASSWORD", "password")
	os.Setenv("DB_NAME", "testdb")
	os.Setenv("DB_SSL_MODE", "disable")
	os.Setenv("DB_MAX_CONNECTIONS", "10")
	os.Setenv("DB_MAX_IDLE_CONNECTIONS", "5")
	os.Setenv("DB_MAX_LIFETIME_CONNECTIONS", "300")
	os.Setenv("REDIS_HOST", "localhost")
	os.Setenv("REDIS_PORT", "6379")
	os.Setenv("REDIS_PASSWORD", "redis_password")
	os.Setenv("REDIS_DB_NUMBER", "0")
	os.Setenv("WEBHOOK_API_URL", "http://example.com/api")
	os.Setenv("WEBHOOK_API_KEY", "api_key")
	os.Setenv("CRON_JOB_SCHEDULE", "0 * * * *") // Every hour
}

func unSetEnvVars() {
	os.Unsetenv("DB_HOST")
	os.Unsetenv("DB_PORT")
	os.Unsetenv("DB_USER")
	os.Unsetenv("DB_PASSWORD")
	os.Unsetenv("DB_NAME")
	os.Unsetenv("DB_SSL_MODE")
	os.Unsetenv("DB_MAX_CONNECTIONS")
	os.Unsetenv("DB_MAX_IDLE_CONNECTIONS")
	os.Unsetenv("DB_MAX_LIFETIME_CONNECTIONS")
	os.Unsetenv("REDIS_HOST")
	os.Unsetenv("REDIS_PORT")
	os.Unsetenv("REDIS_PASSWORD")
	os.Unsetenv("REDIS_DB_NUMBER")
	os.Unsetenv("WEBHOOK_API_URL")
	os.Unsetenv("WEBHOOK_API_KEY")
	os.Unsetenv("CRON_JOB_SCHEDULE")
}

func TestParseEnvironmentVariables(t *testing.T) {

	mockEnvironmentVariables()

	// Clean up environment variables after the test
	defer unSetEnvVars()

	// Call the function being tested
	err := ParseEnvironmentVariables()

	// Check if there's any error
	if err != nil {
		t.Fatalf("ParseEnvironmentVariables failed: %v", err)
	}

	// Verify that environment variables are correctly parsed
	expectedVars := Environment{
		DatabaseHost:            "localhost",
		DatabasePort:            "5432",
		DatabaseUser:            "user",
		DatabasePassword:        "password",
		DatabaseName:            "testdb",
		DatabaseSslMode:         "disable",
		DatabaseMaxConn:         10,
		DatabaseMaxIdleConn:     5,
		DatabaseMaxLifetimeConn: 300,
		RedisHost:               "localhost",
		RedisPort:               "6379",
		RedisPassword:           "redis_password",
		RedisDbNumber:           0,
		WebhookApiUrl:           "http://example.com/api",
		WebhookApiKey:           "api_key",
		CronJobSchedule:         "0 * * * *",
	}

	if EnvironmentVars != expectedVars {
		t.Errorf("Parsed environment variables do not match expected: got %+v, want %+v", EnvironmentVars, expectedVars)
	}
}
