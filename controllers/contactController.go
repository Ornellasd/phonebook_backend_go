package controllers

import (
	"context"
	"log"
	"phonebook_rest_api/config"
	"phonebook_rest_api/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllContacts(c *fiber.Ctx) error {
	contactCollection := config.MI.DB.Collection("entries")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := contactCollection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	var contacts []models.Entry

	if err = cursor.All(ctx, &contacts); err != nil {
		log.Fatal(err)
	}

	return c.Status(fiber.StatusOK).JSON(
		contacts,
	)
}
