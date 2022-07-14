package controller

import (
	"context"
	"strconv"
	"workspace_booking/migration"
	"workspace_booking/model"
	"workspace_booking/utility"

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

func GetAvaliableBookingSpace(c *fiber.Ctx) error {
	// query
	floor_id := c.Query("floor_id")
	from_date := c.Query("from_date")
	to_date := c.Query("to_date")

	// Type Casting
	floorId, _ := strconv.Atoi(floor_id)
	var total_work_space *int

	// DB query call.
	rows := migration.DbPool.QueryRow(context.Background(),
		"select SUM(workspaces_booked) as total_workspace from bookings where floor_id = $1 and from_date >= $2 and to_date <= $3", floorId, from_date, to_date)

	// getting total number of booking space
	floor := model.GetFloorByID(floorId)

	rows.Scan(&total_work_space)
	available_work_space := (floor.TotalWorkSpace - *total_work_space)
	if err := c.JSON(&fiber.Map{
		"success":             true,
		"available_workspace": available_work_space,
	}); err != nil {
		return utility.ErrResponse(c, "Error in getting buildings", 500, err)
	}
	return nil
}

func WorkSpacesDetails(c *fiber.Ctx) error {

	workspace_details := model.GetAllDetails()
	if err := c.JSON(&fiber.Map{
		"success":           true,
		"workspace_details": workspace_details,
		"message":           "All workspace details returned successfully",
	}); err != nil {
		return utility.ErrResponse(c, "Error in getting workspace details", 500, err)
	}
	return nil
}

func MyBookingDetails(c *fiber.Ctx) error {
	userid_params := c.Query("user_id")
	userid, _ := strconv.Atoi(userid_params)
	workspace_details := model.GetMyBookingDetails(true, userid)
	past_booking_details := model.GetMyBookingDetails(false, userid)
	if err := c.JSON(&fiber.Map{
		"success":                   true,
		"upcomming_booking_details": workspace_details,
		"past_booking_details":      past_booking_details,
		"message":                   "All My bookings returned successfully",
	}); err != nil {
		return utility.ErrResponse(c, "Error in getting My bookings", 500, err)
	}
	return nil
}
