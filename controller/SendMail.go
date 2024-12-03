package controller

import (
	"gosendmail/helper"
	"gosendmail/model"

	"github.com/gofiber/fiber/v2"
)

func SendMessage(c *fiber.Ctx) error {

	var payloadSendEmail model.PayloadSendEmail
	if err := c.BodyParser(&payloadSendEmail); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid payload request !",
		})
	}

	helper.SendEmail(payloadSendEmail.To, payloadSendEmail.Subject, payloadSendEmail.Body)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": fiber.StatusOK,
		"data":   payloadSendEmail,
	})
}
