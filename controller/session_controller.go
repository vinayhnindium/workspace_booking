package controller

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"workspace_booking/config"
	m "workspace_booking/model"

	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	log.Println(data)

	log.Println([]byte(data["password"]))

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 10)

	u := &m.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}
	log.Println("pss", password)

	err := u.InsertUser()
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	log.Println("u.pss", u.Password)
	log.Println(u)
	return c.JSON(&fiber.Map{
		"message": "Successfully register",
		"user":    u,
	})
}

func Login(c *fiber.Ctx) error {
	// user := c.FormValue("user")
	// pass := c.FormValue("pass")

	// Throws Unauthorized error
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	u := new(m.User)

	if err := c.BodyParser(u); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	err := u.LoginUser()

	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if u.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	err = bcrypt.CompareHashAndPassword(u.Password, []byte(data["password"]))

	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"id":    u.ID,
		"name":  u.Name,
		"email": u.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.GetJWTSecret()))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Cookie(&fiber.Cookie{
		Name:     u.Email,
		Value:    t,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		SameSite: "lax",
	})

	return c.JSON(fiber.Map{
		"message": "Sucessfully login",
		"token":   t,
		"user":    u,
	})
}

func Logout(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	log.Println(email)
	c.Cookie(&fiber.Cookie{
		Name:     email,
		Value:    "",
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		SameSite: "lax",
	})
	return c.JSON(fiber.Map{
		"message": "Successfully logout",
	})
}
