package routes

import (
	"phonebook_rest_api/controllers"

	"github.com/gofiber/fiber/v2"
)

func ConcactsRoute(route fiber.Router) {
	route.Get("/", controllers.GetAllContacts)
}
