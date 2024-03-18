package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// Ping func for describe ping route.
//
//	@Description	Check if the server is alive
//	@Summary		Check if the server is alive
//	@Tags			Ping
//	@Accept			*/*
//	@Produce		json
//	@Success		200	{string}	string	"Returns PONG if the server is alive"
//	@Router			/ping [get]
func Ping(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "PONG",
	})
}
