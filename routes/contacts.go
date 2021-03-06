package routes

import (
	"phonebook_rest_api/controllers"

	"github.com/gofiber/fiber/v2"
)

func ContactsRoute(route fiber.Router) {
	route.Get("/", controllers.GetAllContacts)
	route.Get("/:id", controllers.GetContact)
	route.Post("/", controllers.AddContact)
	route.Delete("/:id", controllers.DeleteContact)
	route.Put("/:id", controllers.UpdateContact)
}
