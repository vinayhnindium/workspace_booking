package controller

import (
	"log"
	"workspace_booking/model"

	"github.com/gofiber/fiber/v2"
)

// Index
func AllLocations(c *fiber.Ctx) error {
	locations := model.GetAllLocations()
	if len(locations) != 0 {
		if err := c.JSON(&fiber.Map{
			"success":   true,
			"locations": locations,
			"message":   "All Loations returned successfully",
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
			"message": "No Records found for Location",
		})
	}
	return nil
}

// CreateLocation handler
func CreateLocation(c *fiber.Ctx) error {
	location := new(model.Location)
	if err := c.BodyParser(location); err != nil {
		log.Println(err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	err := location.CreateLocation()
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}

	// Return Location in JSON format
	if location.Id != 0 {
		if err := c.JSON(&fiber.Map{
			"success":  true,
			"location": location,
			"message":  "Location successfully created",
		}); err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"message": "Error creating Location",
			})
		}
	} else {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Location table not found",
		})
	}
	return nil
}
