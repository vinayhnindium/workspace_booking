package controller

import (
	"context"
	"fmt"
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

	if err := c.JSON(&fiber.Map{
		"success": true,
		"role":    workspaceParams,
		"message": "Booking successfully created",
	}); err != nil {
		return utility.ErrResponse(c, "Error in response", 500, err)
	}
	return nil
}

func GetAvaliableBookingSpace(c *fiber.Ctx)error{
// query 
floor_id := c.Query("floor_id")
from_date := c.Query("from_date")
to_date := c.Query("to_date")

// DB query call.
rows := migration.DbPool.QueryRow(context.Background(), 
"select SUM(workspaces_booked) from bookings where floor_id=$1 and from_date=$2 and to_date=$3", floor_id, from_date, to_date)

intVar, _ := strconv.Atoi(floor_id)

floor := model.GetFloorByID(intVar)

fmt.Println(rows)
fmt.Println(floor.TotalWorkSpace)

if err := c.JSON(&fiber.Map{
	"success":   true,
	"available_workspace": rows,
	"message":   "All Buildings returned successfully",
}); err != nil {
	return utility.ErrResponse(c, "Error in getting buildings", 500, err)
}


return nil
}
