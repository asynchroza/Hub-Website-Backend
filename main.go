package main

import (
	"hub-backend/configs"
	"hub-backend/routes"

	"github.com/gofiber/fiber/v2"
	// fmt, encoding/json, strconv
)

func main() {

	app := fiber.New()
	configs.ConnectDB()
	routes.MembersRoute(app)
	app.Listen(":6000")
}
