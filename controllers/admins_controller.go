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
	err := adminCollection.FindOne(ctx, bson.M{"user": incomming_admin.Username}).Decode(&user)

	// fmt.Println(user)

	if err != nil {
		fmt.Println(fiber.StatusBadRequest)
	} else {
		fmt.Println(fiber.StatusOK)
	}

	p, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(p))

	if err := bcrypt.CompareHashAndPassword(p, []byte(incomming_admin.Password)); err != nil {
		fmt.Println(incomming_admin.Password)
		return c.JSON(fiber.Map{
			"Message": err.Error(),
		})
	}

	return c.JSON(user)
}
