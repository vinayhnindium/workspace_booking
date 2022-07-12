package controller

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"workspace_booking/model"
	"workspace_booking/utility"
)

// CreateBooking handler
func CreateBooking(c *fiber.Ctx) error {
	workspaceParams := new(model.Booking)

	if err := c.BodyParser(workspaceParams); err != nil {
		log.Println(err)
		return utility.ErrResponse(c, "Error in body parsing", 400)
	}
	err := workspaceParams.InsertBooking()

	if err != nil {
		log.Println(err)
		return utility.ErrResponse(c, "Error in creation", 500)
	}

	if err := c.JSON(&fiber.Map{
		"success": true,
		"role":    workspaceParams,
		"message": "Booking successfully created",
	}); err != nil {
		return utility.ErrResponse(c, "Error in response", 500)
	}
	return nil
}
