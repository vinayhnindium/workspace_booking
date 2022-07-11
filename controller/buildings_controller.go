package controller

import (
	"log"
	"workspace_booking/database"
	"workspace_booking/model"

	"github.com/gofiber/fiber/v2"
)

// Index
func Index(c *fiber.Ctx) error {
	// query all data
	rows, e := database.DB.Query("select * from buildings")
	if e != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Error occurred in Building module",
		})
	}
	defer rows.Close()

	// declare empty post variable
	result := model.Buildings{}

	// iterate over rows
	for rows.Next() {
		building := model.Building{}
		e = rows.Scan(&building.Id, &building.Name, &building.City, &building.Area, &building.Address, &building.UpdatedAt, &building.CreatedAt)
		if e != nil {
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"message": "Error occurred in Building module",
			})
		}
		result.Buildings = append(result.Buildings, building)
	}
	return c.JSON(&fiber.Map{
		"success":  true,
		"building": result,
		"message":  "All building returned successfully",
	})
}

// to create new building record
func CreateBuilding(c *fiber.Ctx) error {
	// Instantiate new Building struct
	r := new(model.Building)
	//  Parse body into building struct
	if err := c.BodyParser(r); err != nil {
		log.Println(err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	log.Println(r)
	// Insert Building into database
	sqlStatement := "INSERT INTO buildings (name, city, area, address) VALUES ($1, $2, $3, $4)"
	err := database.DB.QueryRow(sqlStatement, r.Name, r.City, r.Area, r.Address)
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}

	// Return Building in JSON format
	if err := c.JSON(&fiber.Map{
		"success":  true,
		"message":  "Building successfully created",
		"building": r,
	}); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": "Error occurred in Building module",
		})
	}
	return nil
}

// Update
func Update(c *fiber.Ctx) error {
	r := new(model.Building)
	//  Parse body into building struct
	if err := c.BodyParser(&r); err != nil {
		log.Println(err)
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	//Update db
	sqlStatement := "UPDATE buildings set name=$1, area=$2, city=$3, address=$4 WHERE id=$5"
	database.DB.Exec(sqlStatement, r.Name, r.Area, r.City, r.Address, r.Id)
	building := model.Building{}
	row := database.DB.QueryRow("SELECT * FROM buildings WHERE id = $1", r.Id)
	row.Scan(&building.Id, &building.Name, &building.City, &building.Area, &building.Address, &building.UpdatedAt, &building.CreatedAt)
	return c.Status(200).JSON(&fiber.Map{
		"success":  true,
		"building": building,
		"message":  "Building successfully Updated",
	})
}

// Show
func Show(c *fiber.Ctx) error {
	building := model.Building{}
	id := c.Params("id")
	row := database.DB.QueryRow("SELECT * FROM buildings WHERE id = $1", id)
	row.Scan(&building.Id, &building.Name, &building.City, &building.Area, &building.Address, &building.UpdatedAt, &building.CreatedAt)
	if building.Id != 0 {
		return c.Status(200).JSON(&fiber.Map{
			"success":  true,
			"building": building,
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"message": "Record not found",
	})
}

// Delete
func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	row := database.DB.QueryRow("SELECT * FROM buildings WHERE id = $1", id)
	if row != nil {
		sqlStatement := "delete FROM buildings WHERE id = $1"
		database.DB.Exec(sqlStatement, id)
		return c.Status(200).JSON(&fiber.Map{
			"success": true,
			"message": "Building Record deleted successfully",
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"message": "Record not found",
	})
}
