package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/root/medicalBooking/database"
	"github.com/root/medicalBooking/models"
)

func CreateUser(c *fiber.Ctx) error {
	context := fiber.Map{
		"status":  "success",
		"message": "Create a user",
	}
	var newUser models.User
	if err := c.BodyParser(&newUser); err != nil {
		log.Panic("Error in parsing user")
	}
	database.DBConn.Create(&newUser)
	context["data"] = newUser

	return c.Status(200).JSON(context)
}

func GetAllUsers(c *fiber.Ctx) error {
	context := fiber.Map{
		"status":  "success",
		"message": "Get all users",
	}
	var getAllUsers []models.User
	database.DBConn.Find(&getAllUsers)
	context["data"] = getAllUsers

	return c.Status(200).JSON(context)
}

func GetSingleUser(c *fiber.Ctx) error {
	context := fiber.Map{
		"status":  "success",
		"message": "Get all users",
	}
	id := c.Params("id")
	var getSingleUser models.User
	database.DBConn.First(&getSingleUser, id)
	// if getSingleUser.ID == 0 {
	// 	log.Panic("Enter an ideal number")
	// }
	if err := c.BodyParser(&getSingleUser); err != nil {
		log.Panic("User does not exit")
	}
	context["data"] = getSingleUser

	return c.Status(200).JSON(context)
}

func UpdateUser(c *fiber.Ctx) error {
	context := fiber.Map{
		"status":  "success",
		"message": "Update user",
	}
	id := c.Params("id")
	var updateUser models.User
	database.DBConn.First(&updateUser, id)
	if err := c.BodyParser(&updateUser); err != nil {
		log.Panic("Error in updating user")
	}
	database.DBConn.Save(&updateUser)
	context["data"] = updateUser

	return c.Status(200).JSON(context)
}

func DeleteUser(c *fiber.Ctx) error {
	context := fiber.Map{
		"status":  "success",
		"message": "Delete a user",
	}
	id := c.Params("id")
	var delUser models.User
	database.DBConn.First(&delUser, id)
	if err := c.BodyParser(&delUser); err != nil {
		log.Panic("Error in deleting user")
	}
	database.DBConn.Delete(&delUser)
	context["data"] = "User has been deleted..."
	context["data"] = delUser

	return c.Status(200).JSON(context)
}
