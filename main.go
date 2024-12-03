package main

import (
	"gosendmail/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000", //for example allow origin for reactJS
		AllowCredentials: true,
	}))

	routes.SetUpRoutes(app)

	app.Listen(":8000")
}
