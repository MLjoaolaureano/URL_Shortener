package router

import (
	"awesomeProject1/handler"
	"github.com/gofiber/fiber"
)

type Router interface {
	SetupRoutes(app *fiber.App)
}

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	api.Get("/test", handler.Test)
	api.Post("/generate", handler.CreateShortUrl)
	api.Get("/recover/:url", handler.RecoverUrl)
}
