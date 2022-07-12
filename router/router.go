package router

import (
	"workspace_booking/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	api.Get("/roles", controller.AllRoles)
	api.Post("/roles", controller.CreateRole)

	api.Post("/book_workspace", controller.CreateBooking)

	api.Get("/users", controller.GetUsers)
	api.Post("/users", controller.CreateUser)
	api.Get("/users/:id", controller.GetUser)
	api.Delete("/users/:id", controller.DeleteUser)
}
