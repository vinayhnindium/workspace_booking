package utility

import "github.com/gofiber/fiber/v2"

func ErrResponse(c *fiber.Ctx, message string, status int) error {
	return c.Status(status).JSON(&fiber.Map{
		"success": false,
		"message": message,
	})
}
