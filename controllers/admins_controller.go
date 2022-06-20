package controllers

import (
	"context"
	"fmt"
	"hub-backend/configs"
	"hub-backend/models"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var adminCollection *mongo.Collection = configs.GetCollection(configs.DB, "admins")

func LoginAdmin(c *fiber.Ctx) error {
	var incomming_admin models.Admin

	if err := c.BodyParser(&incomming_admin); err != nil {
		log.Fatal(err.Error())
		return err
	}

	var user models.Admin

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := adminCollection.FindOne(ctx, bson.M{"username": incomming_admin.Username}).Decode(&user)

	fmt.Println(user)

	if err != nil {
		fmt.Println(fiber.StatusBadRequest)
	} else {
		fmt.Println(fiber.StatusOK)
	}

	// p, err := bcrypt.GenerateFromPassword([]byte(incomming_admin.Password), 10)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(incomming_admin.Password)); err != nil {
		log.Fatalf(err.Error())
	}

	return c.JSON(user)
}
