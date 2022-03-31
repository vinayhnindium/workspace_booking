package controller

import (
	"log"
	"workspace_booking/database"
	"workspace_booking/model"

	"github.com/gofiber/fiber/v2"
)

// GetAllRoles from db
func GetAllRoles(c *fiber.Ctx) error {
	// query role table in the database
	rows, err := database.DB.Query("SELECT * FROM roles")
	if err != nil {
		log.Println(1, err)
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	defer rows.Close()
	result := model.Roles{}

	for rows.Next() {
		role := model.Role{}
		err := rows.Scan(&role.Id, &role.Name, &role.UpdatedAt, &role.CreatedAt)
		// Exit if we get an error
		if err != nil {
			log.Println(2, err)
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}
		// Append Role to Roles
		result.Roles = append(result.Roles, role)
	}
	// Return Roles in JSON format
	if err := c.JSON(&fiber.Map{
		"success": true,
		"role":    result,
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
	r := new(model.Role)
	//  Parse body into role struct
	if err := c.BodyParser(r); err != nil {
		log.Println(err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	// Insert Role into database
	res, err := database.DB.Query("INSERT INTO roles (name) VALUES ($1)", r.Name)
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	// Print result
	log.Println(res)

	// Return Role in JSON format
	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "Role successfully created",
		"role":    r,
	}); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Error creating role",
		})
	}
	return nil
}
