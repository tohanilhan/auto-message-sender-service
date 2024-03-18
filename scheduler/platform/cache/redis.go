package cache

import (
	"github.com/redis/go-redis/v9"
	"github.com/tohanilhan/auto-message-sender-service/scheduler/pkg/utils"
)

// RedisConnection func for connect to Redis server.
func RedisConnection() (*redis.Client, error) {
	// Define Redis database number.

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
