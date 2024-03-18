package cache

import (
	"github.com/tohanilhan/auto-message-sender-service/pkg/utils"

	"github.com/redis/go-redis/v9"
)

// RedisConnection func for connect to Redis server.
func RedisConnection() (*redis.Client, error) {

	// Build Redis connection URL.
	redisConnURL, err := utils.ConnectionURLBuilder("redis")
	if err != nil {
		return nil, err
	}

	// Set Redis options.
	options := &redis.Options{
		Addr:     redisConnURL,
		Password: utils.EnvironmentVars.RedisPassword,
		DB:       utils.EnvironmentVars.RedisDbNumber,
	}

	return redis.NewClient(options), nil
}
