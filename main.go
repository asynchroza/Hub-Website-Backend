package main

import (
	"hub-backend/configs"
	"hub-backend/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	// fmt, encoding/json, strconv
)

func main() {

	app := fiber.New()
	configs.ConnectDB()
	routes.MembersRoute(app)
	routes.AdminRoute(app)
	routes.EventsRoute(app)
	log.Fatal(app.Listen(":8000"))
}
