package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/manav-chan/rhapsody/controller"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)

	// app.Use(middleware.IsAuthenticated)
}