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
	api.Get("/workspace_details", controller.WorkSpacesDetails)

	/* Building API's */
	api.Get("/buildings", controller.AllBuildings)
	api.Post("/buildings", controller.CreateBuilding)

	/* City API's */
	api.Get("/cities", controller.AllCities)
	api.Post("/cities", controller.CreateCity)

	/* Location API's */
	api.Get("/locations", controller.AllLocations)
	api.Post("/locations", controller.CreateLocation)

	/* Floor API's */
	api.Get("/floors", controller.AllFloors)
	api.Post("/floors", controller.CreateFloor)

	// Get available work-space

	api.Get("/available_workspace", controller.GetAvaliableBookingSpace)
}
