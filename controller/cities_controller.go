package controller

import (
	"log"
	"workspace_booking/model"

	"github.com/gofiber/fiber/v2"
)

// Index
func AllCities(c *fiber.Ctx) error {
	cities := model.GetAllCities()
	if len(cities) != 0 {
		if err := c.JSON(&fiber.Map{
			"success": true,
			"cities":  cities,
			"message": "All Cities returned successfully",
		}); err != nil {
			log.Println(3, err)
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"message": err,
			})
		}
	} else {
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"message": "No Records found for City",
		})
	}
	return nil
}

// CreateCity handler
func CreateCity(c *fiber.Ctx) error {
	city := new(model.City)
	if err := c.BodyParser(city); err != nil {
		log.Println(err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	err := city.CreateCity()
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}

	// Return city in JSON format
	if err := c.JSON(&fiber.Map{
		"success": true,
		"city":    city,
		"message": "City successfully created",
	}); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Error creating City",
		})
	}
	return nil
}
