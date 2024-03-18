package utils

import (
	"fmt"
)

// ConnectionURLBuilder func for building URL connection.
func ConnectionURLBuilder(n string) (string, error) {
	// Define URL to connection.
	var url string

	// Switch given names.
	switch n {
	case "postgres":
		// URL for PostgreSQL connection.
		url = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			EnvironmentVars.DatabaseHost,
			EnvironmentVars.DatabasePort,
			EnvironmentVars.DatabaseUser,
			EnvironmentVars.DatabasePassword,
			EnvironmentVars.DatabaseName,
			EnvironmentVars.DatabaseSslMode,
		)
	case "redis":
		// URL for Redis connection.
		url = fmt.Sprintf(
			"%s:%s",
			EnvironmentVars.RedisHost,
			EnvironmentVars.RedisPort,
		)
	case "fiber":
		// URL for Fiber connection.
		url = fmt.Sprintf(
			":%s",
			EnvironmentVars.ServerPort,
		)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	// Return connection URL.
	return url, nil
}
