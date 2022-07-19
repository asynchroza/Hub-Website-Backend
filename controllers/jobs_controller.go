package controllers

import (
	"context"
	"hub-backend/configs"
	"hub-backend/models"
	"hub-backend/responses"
	"net/http"
	"reflect"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

var jobsCollection *mongo.Collection = configs.GetCollection(configs.DB, "jobs")

func CreateJob(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	bearer_token := c.Get("BEARER_TOKEN")

	var job models.Job
	defer cancel()

	if bearer_token != configs.ReturnAuthToken() {
		return c.Status(http.StatusBadRequest).JSON(responses.MemberResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"Reason": "Authentication failed"}})
	}

	if err := c.BodyParser(&job); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.MemberResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&job); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.MemberResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newJob := models.Job{
		JobID:       job.JobID,
		Position:    job.Position,
		Company:     job.Company,
		Description: job.Description,
		Link:        job.Link,
	}

	// empty_body_check := EmptyStringBody(newJob.JobID)

	v := reflect.ValueOf(newJob)
	type_of_v := v.Type()

	// this will only work if all struct fields are string
	for i := 0; i < v.NumField(); i++ {

		empty_body_check := EmptyStringBody(v.Field(i).Interface().(string))
		if empty_body_check {
			return isEmptyException(c, type_of_v.Field(i).Name)
		}
	}

	result, err := jobsCollection.InsertOne(ctx, newJob)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.MemberResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.MemberResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})

}
