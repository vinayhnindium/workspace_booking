package controller

import (
	"log"

	database "workspace_booking/model"

	"github.com/gofiber/fiber/v2"
)

// AllRoles from db
func AllRoles(c *fiber.Ctx) error {
	// query role table in the database

	roles := database.GetAllRoles()

	if err := c.JSON(&fiber.Map{
		"success": true,
		"role":    roles,
		"message": "All role returned successfully",
	}); err != nil {
		log.Println(3, err)
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	return nil
}

// CreateRole handler
func CreateRole(c *fiber.Ctx) error {
	// Instantiate new Role struct
	r := new(database.Role)

	//  Parse body into role struct
	if err := c.BodyParser(r); err != nil {
		log.Println(err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	err := r.CreateRole(r.Name)

	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}

	// Print result
	log.Println(r)

	// Return Role in JSON format
	if err := c.JSON(&fiber.Map{
		"success": true,
		"role":    r,
		"message": "Role successfully created",
	}); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Error creating role",
		})
	}
	return nil
}
