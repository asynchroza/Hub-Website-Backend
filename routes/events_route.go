package routes

import (
	"hub-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func EventsRoute(app *fiber.App) {
	app.Put("/api/event", controllers.CreateEvent)
}
