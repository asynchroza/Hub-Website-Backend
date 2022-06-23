package routes

import (
	"hub-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func EventsRoute(app *fiber.App) {
	app.Post("/api/event", controllers.CreateEvent)
	app.Get("/api/event/:key", controllers.GetEvent)
}
