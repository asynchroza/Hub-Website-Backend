package routes

import (
	"hub-backend/controllers"

	"github.com/gofiber/fiber/v2"
)

func MembersRoute(app *fiber.App) {
	// club members requests routes
	app.Post("/api/member", controllers.CreateMember)
	app.Get("/api/members", controllers.GetAllMembers)
}
