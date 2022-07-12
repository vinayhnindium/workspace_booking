package utility

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func ErrResponse(c *fiber.Ctx, message string, status int, err error) error {
	if err != nil {
		log.Println(err)
	}
	return c.Status(status).JSON(&fiber.Map{
		"success": false,
		"message": message,
	})
}
