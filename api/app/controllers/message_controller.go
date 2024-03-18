package controllers

import (
	"context"
	"sort"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/tohanilhan/auto-message-sender-service/app/models"
	"github.com/tohanilhan/auto-message-sender-service/platform/cache"
)

// GetMessages func for get sent messages from Redis.
//
//	@Description	Can be used to get all sent messages from Redis.
//	@Summary		Get all sent messages from Redis.
//	@Tags			Messages
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	models.Message
//	@Router			/messages [get]
//	@Security		ApiKeyAuth
func GetMessages(c *fiber.Ctx) error {

	// Get Redis connection
	redisConn, err := cache.RedisConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Error opening Redis connection.",
		})
	}

	// Get keys from Redis
	keys, err := redisConn.Keys(context.Background(), "sent_message:*").Result()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Error getting messages from Redis.",
		})
	}

	// Get messages from Redis
	messages := []models.Message{}
	for _, key := range keys {
		value, err := redisConn.Get(context.Background(), key).Result()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": "Error getting messages from Redis.",
			})
		}

		message := models.Message{
			MessageID: strings.Replace(key, "sent_message:", "", 1),
			SentTime:  value,
		}

		messages = append(messages, message)
	}

	sortedMessageList := sortMessages(messages)

	// Return status 200 OK.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "Messages retrieved successfully.",
		"data": fiber.Map{
			"count":    len(messages),
			"messages": sortedMessageList,
		},
	})
}

func sortMessages(messages []models.Message) []models.Message {
	// Sort messages.
	sort.Slice(messages, func(i, j int) bool {
		return messages[i].SentTime < messages[j].SentTime
	})

	return messages
}
