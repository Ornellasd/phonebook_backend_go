package controllers

import (
	"context"
	"fmt"
	"log"
	"phonebook_rest_api/config"
	"phonebook_rest_api/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllContacts(c *fiber.Ctx) error {
	contactCollection := config.MI.DB.Collection("entries")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := contactCollection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(ctx)

	var contacts []models.Entry

	if err = cursor.All(ctx, &contacts); err != nil {
		log.Fatal(err)
	}

	return c.Status(fiber.StatusOK).JSON(
		contacts,
	)
}

func GetContact(c *fiber.Ctx) error {
	contactCollection := config.MI.DB.Collection("entries")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	var contact models.Entry
	objId, err := primitive.ObjectIDFromHex(c.Params("id"))
	findResult := contactCollection.FindOne(ctx, bson.M{"_id": objId})
	if err := findResult.Err(); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Catchphrase Not found",
			"error":   err,
		})
	}

	err = findResult.Decode(&contact)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Catchphrase Not found",
			"error":   err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(
		contact,
	)
}

func AddContact(c *fiber.Ctx) error {
	contactCollection := config.MI.DB.Collection("entries")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	contact := new(models.Entry)

	if err := c.BodyParser(contact); err != nil {
		log.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	result, err := contactCollection.InsertOne(ctx, contact)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success":  false,
			"messsage": "Catchphrase failed to insert",
			"error":    err,
		})
	}

	fmt.Println(result)

	return c.Status(fiber.StatusCreated).JSON(
		result,
	)
}
