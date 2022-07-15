package controller

import (
	"context"
	"fmt"
	"strconv"
	"workspace_booking/migration"
	"workspace_booking/model"
	"workspace_booking/utility"

	"github.com/golang-jwt/jwt/v4"

	"github.com/gofiber/fiber/v2"
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

func GetAvailableBookingSpace(c *fiber.Ctx) error {

	type params struct {
		FloorId  int    `json:"floor_id"`
		FromDate string `json:"from_date"`
		ToDate   string `json:"to_date"`
	}

	workspaceParams := new(params)

	if err := c.BodyParser(workspaceParams); err != nil {
		return utility.ErrResponse(c, "Error in body parsing", 400, err)
	}

	var totalWorkSpace *int

	// DB query call.
	rows := migration.DbPool.QueryRow(context.Background(),
		"select SUM(workspaces_booked) as total_workspace from bookings where floor_id = $1 and from_date >= $2 and to_date <= $3", workspaceParams.FloorId, workspaceParams.FromDate, workspaceParams.ToDate)

	// getting total number of booking space
	floor := model.GetFloorByID(workspaceParams.FloorId)

	err := rows.Scan(&totalWorkSpace)
	if err != nil {
		return err
	}
	availableWorkSpace := floor.TotalWorkSpace - *totalWorkSpace
	if err := c.JSON(&fiber.Map{
		"success":             true,
		"available_workspace": availableWorkSpace,
	}); err != nil {
		return utility.ErrResponse(c, "Error in getting buildings", 500, err)
	}
	return nil
}

func WorkSpacesDetails(c *fiber.Ctx) error {

	workspaceDetails := model.GetAllDetails()
	if err := c.JSON(&fiber.Map{
		"success":           true,
		"workspace_details": workspaceDetails,
		"message":           "All workspace details returned successfully",
	}); err != nil {
		return utility.ErrResponse(c, "Error in getting workspace details", 500, err)
	}
	return nil
}

func MyBookingDetails(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	var userId int

	currentUserId := fmt.Sprintf("%v", claims["id"])
	userId, _ = strconv.Atoi(currentUserId)

	workspaceDetails := model.GetMyBookingDetails(true, userId)
	pastBookingDetails := model.GetMyBookingDetails(false, userId)
	if err := c.JSON(&fiber.Map{
		"success":                  true,
		"upcoming_booking_details": workspaceDetails,
		"past_booking_details":     pastBookingDetails,
		"message":                  "All My bookings returned successfully",
	}); err != nil {
		return utility.ErrResponse(c, "Error in getting My bookings", 500, err)
	}
	return nil
}
