package main

import (
	"context"
	"fmt"
	"log"
	"workspace_booking/config"
	db "workspace_booking/database"
	"workspace_booking/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	dbPool := db.GetDbConnectionPool()
	// close database
	defer dbPool.Close()
	// check db
	err := dbPool.Ping(context.Background())
	CheckError(err)

	fmt.Println("Connected!")

	app := fiber.New()
	app.Use(logger.New())
	router.SetupRoutes(app)

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	u := db.User{}
	// 	users := u.GetUsers()
	// 	fmt.Println(users)
	// 	return c.JSON(users)
	// })

	println(config.GetServerPort())
	log.Fatalln(app.Listen(config.GetServerPort()))
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
