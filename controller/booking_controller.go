package controller

import (
	"context"
	"strconv"
	"workspace_booking/config"
	"workspace_booking/migration"
	"workspace_booking/model"
	"workspace_booking/utility"

	"github.com/gofiber/fiber/v2"
)

// CreateBooking handler
func CreateBooking(c *fiber.Ctx) error {
	timingParams := new(model.BookingTiming)

	c.BodyParser(timingParams)

	fromDatetTime, toDateTime := model.BookingTimestamp(timingParams)

	workspaceParams := new(model.Booking)

	if err := c.BodyParser(workspaceParams); err != nil {
		return utility.ErrResponse(c, "Error in body parsing", 400, err)
	}

	workspaceParams.FromDateTime = fromDatetTime
	workspaceParams.ToDateTime = toDateTime

	err := workspaceParams.InsertBooking()

	if err != nil {
		return utility.ErrResponse(c, "Error in creation", 500, err)
	}

	err = model.BulkInsertBookingParticipant(workspaceParams)

	err = model.BulkInsertBookingWorkspace(workspaceParams, timingParams)

	if err != nil {
		return utility.ErrResponse(c, "Error in creating participants", 500, err)
	}

	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "Booking successfully created",
	}); err != nil {
		return utility.ErrResponse(c, "Error in response", 500, err)
	}
	return nil
}

func GetAvailableBookingSpace(c *fiber.Ctx) error {
	reqFloorId := c.Query("floor_id")
	fromDate := c.Query("from_datetime")
	toDate := c.Query("to_datetime")
	floorId, _ := strconv.Atoi(reqFloorId)

	var totalWorkSpace *int

	// DB query call.
	rows := migration.DbPool.QueryRow(context.Background(),
		"select SUM(workspaces_booked) as total_workspace from bookings where floor_id = $1 and from_datetime >= $2 and to_datetime <= $3", floorId, fromDate, toDate)

	// getting total number of booking space
	floor := model.GetFloorByID(floorId)

	err := rows.Scan(&totalWorkSpace)
	if err != nil {
		return err
	}
	var availableWorkSpace int

	if availableWorkSpace = 0; totalWorkSpace != nil {
		availableWorkSpace = floor.TotalWorkSpace - *totalWorkSpace
	} else {
		availableWorkSpace = floor.TotalWorkSpace
	}

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
	auth, err := config.GetAuthDetails(c)
	if err != nil {
		return utility.ErrResponse(c, "Error in getting buildings", 500, err)
	}

	var userId int

	userId, _ = strconv.Atoi(auth.UserID)

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
