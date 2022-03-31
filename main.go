package main

import (
	// "database/sql"

	"workspace_booking/database"
	"workspace_booking/router"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

type SendJsonResponse struct {
	Message string
}

func main() {

	database.Connect()

	app := fiber.New()
	router.SetupRoutes(app)

	app.Listen(":3000")
}
