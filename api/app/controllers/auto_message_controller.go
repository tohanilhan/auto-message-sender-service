package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tohanilhan/auto-message-sender-service/pkg/repository"
	"github.com/tohanilhan/auto-message-sender-service/platform/database"
)

// ChangeAutoSendingBehaviour func for toggle auto sending messages.
//
//	@Description	This is the endpoint for changing the behaviour of the message sending service.
//	@Summary		Can be used to start/stop auto messages sending feature.
//	@Tags			Config
//	@Accept			*/*
//	@Produce		json
//	@Param			option	path		string	true	"Option to on/off auto sending messages"	Enums(ON, OFF)
//	@Success		200		{string}	string	"Returns the current status of the auto sending feature."
//	@Router			/change-auto-sending/{option} [post]
//	@Security		ApiKeyAuth
func ChangeAutoSendingBehaviour(c *fiber.Ctx) error {

	option := c.Params("option")

	if option != repository.AutoSendOn && option != repository.AutoSendOff {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": fmt.Sprintf("Invalid option. Please use %s or %s.", repository.AutoSendOn, repository.AutoSendOff),
		})
	}

	db, err := database.OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Error opening database connection.",
		})
	}

	err = db.UpdateAutoSendingConfig(option)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Error updating auto sending config.",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": fmt.Sprintf("Auto sending feature is now %s.", option),
	})
}
