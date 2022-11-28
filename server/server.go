package server

import (
	"awesomeProject1/config"
	"awesomeProject1/router"
	"github.com/gofiber/fiber"
)

func Setup() {
	app := fiber.New()

	router.SetupRoutes(app)
	app.Listen(config.Config("PORT"))
}
