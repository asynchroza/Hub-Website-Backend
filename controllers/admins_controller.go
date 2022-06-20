package controllers

import (
	"hub-backend/configs"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

var adminCollection *mongo.Collection = configs.GetCollection(configs.DB, "admins")

func GetAdmin(c *fiber.Ctx) error {
	// return errors
}
