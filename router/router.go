package router

import (
	"workspace_booking/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	api.Get("/roles", controller.AllRoles)
	api.Post("/roles", controller.CreateRole)

}
