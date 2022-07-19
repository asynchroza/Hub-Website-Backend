package routes

import (
	"hub-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func JobsRoute(app *fiber.App) {
	// validate Admin
	app.Post("/api/jobs", controllers.CreateJob)
}
