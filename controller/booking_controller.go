package controller

import (
	"github.com/gofiber/fiber/v2"
	"workspace_booking/model"
	"workspace_booking/utility"
)

// CreateBooking handler
func CreateBooking(c *fiber.Ctx) error {
	workspaceParams := new(model.Booking)

	if err := c.BodyParser(workspaceParams); err != nil {
		return utility.ErrResponse(c, "Error in body parsing", 400, err)
	}
	err := workspaceParams.InsertBooking()

	if err != nil {
		return utility.ErrResponse(c, "Error in creation", 500, err)
	}

	err = model.BulkInsertBookingParticipant(workspaceParams.Id, workspaceParams.UserIds)

	if err != nil {
		return utility.ErrResponse(c, "Error in creating participants", 500, err)
	}

	if err := c.JSON(&fiber.Map{
		"success": true,
		"role":    workspaceParams,
		"message": "Booking successfully created",
	}); err != nil {
		return utility.ErrResponse(c, "Error in response", 500, err)
	}
	return nil
}
