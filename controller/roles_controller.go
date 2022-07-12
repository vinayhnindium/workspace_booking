package controller

import (
	"log"
	"workspace_booking/utility"

	"workspace_booking/model"

	"github.com/gofiber/fiber/v2"
)

// AllRoles from db
func AllRoles(c *fiber.Ctx) error {
	// query role table in the model

	roles := model.GetAllRoles()

	if err := c.JSON(&fiber.Map{
		"success": true,
		"role":    roles,
		"message": "All role returned successfully",
	}); err != nil {
		log.Println(3, err)
		return utility.ErrResponse(c, "Error in getting roles", 500)
	}
	return nil
}

// CreateRole handler
func CreateRole(c *fiber.Ctx) error {
	// Instantiate new Role struct
	r := new(model.Role)

	//  Parse body into role struct
	if err := c.BodyParser(r); err != nil {
		log.Println(err)
		return utility.ErrResponse(c, "Error in parsing", 400)
	}
	err := r.InsertRole()

	if err != nil {
		log.Println(err)
		return utility.ErrResponse(c, "Error in saving", 500)
	}

	// Print result
	log.Println(r)

	// Return Role in JSON format
	if err := c.JSON(&fiber.Map{
		"success": true,
		"role":    r,
		"message": "Role successfully created",
	}); err != nil {
		return utility.ErrResponse(c, "Error in sending response", 500)
	}
	return nil
}
