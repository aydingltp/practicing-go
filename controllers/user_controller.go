package controllers

import (
	"github.com/gofiber/fiber/v2"
	"practicing-go/database"
	"practicing-go/models"
)

var db = database.DB

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	db.Find(&user, id)
	if user.Name == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with Id", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "User found", "data": user})
}

func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User
	db.Find(&users)
	return c.JSON(fiber.Map{"status": "success", "message": "All users", "data": users})
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{})
	}

	db.Create(user)
	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": user})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	db.First(&user, id)

	if user.Name == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
	}
	db.Delete(&user)
	return c.JSON(fiber.Map{"status": "success", "message": "User successfully deleted", "data": nil})
}
