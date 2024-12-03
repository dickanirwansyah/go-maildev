package routes

import (
	"gosendmail/controller"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	app.Post("/api/send-email", controller.SendMessage)
}
