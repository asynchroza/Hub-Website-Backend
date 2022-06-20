package routes

import (
	"hub-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func AdminRoute(app *fiber.App) {
	// validate Admin
	app.Post("/api/admin", controllers.GetAdmin)
}
