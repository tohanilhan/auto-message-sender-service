package middleware

import (
	"crypto/sha256"
	"crypto/subtle"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"github.com/tohanilhan/auto-message-sender-service/pkg/utils"
)

func KeyAuthProtected() fiber.Handler {
	config := keyauth.Config{
		KeyLookup: "header:x-ins-api-auth-key",
		Validator: ValidateAPIKey,
	}
	return keyauth.New(config)
}

func ValidateAPIKey(c *fiber.Ctx, key string) (bool, error) {

	// Compare the provided key with the actual API key.
	// The subtle package is used to avoid timing attacks.
	hashedAPIKey := sha256.Sum256([]byte(strings.TrimSpace(utils.EnvironmentVars.ApiKey))) // Hash the API key.
	hashedKey := sha256.Sum256([]byte(strings.TrimSpace(key)))

	if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
		return true, nil
	}

	return false, keyauth.ErrMissingOrMalformedAPIKey
}
