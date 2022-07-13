package controller

import (
	"strconv"
	m "workspace_booking/model"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	u := new(m.User)

	if err := c.BodyParser(u); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	err := u.InsertUser()
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"message": "User Successfully Added",
		"user":    u,
	})
}

func GetUsers(c *fiber.Ctx) error {
	users := new(m.Users)

	users, err := users.FetchUsers()

	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"users": users,
	})
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	i, e := strconv.Atoi(id)

	if e != nil {
		return c.Status(400).SendString(e.Error())
	}

	u := m.User{ID: i}
	err := u.FetchUser()

	// user, err := m.FetchUser(db, id)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"user": u,
	})
}

func EditUser(c *fiber.Ctx) error {
	id := c.Params("id")

	i, e := strconv.Atoi(id)

	if e != nil {
		return c.Status(400).SendString(e.Error())
	}
	u := &m.User{ID: i}

	if err := c.BodyParser(u); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	err := u.UpdateUser()

	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"message": "User Successfully Updated",
		"user":    u,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	i, e := strconv.Atoi(id)

	if e != nil {
		return c.Status(400).SendString(e.Error())
	}

	u := &m.User{ID: i}

	err := u.DeleteUser()

	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"message": "User Successfully Deleted",
	})
}
