package controllers

import (
	"log"

	"github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/db"
	m "github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreatePermission(c *fiber.Ctx) error {

	// Declare a variable to hold the request body
	var data map[string]string

	// Parse the request body into the data variable
	err := c.BodyParser(&data)
	if err != nil {
		// If there's an error in parsing the body, log the error and return an Internal Server Error response
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}

	var permission *m.Permission
	result := db.DB.Where(&m.Permission{Name: data["name"]}).Find(&permission)
	if result.RowsAffected == 1 {
		return c.Status(fiber.StatusOK).JSON(&m.Response{Success: false, Message: "Permission already exists"})
	}

	permission = &m.Permission{
		ID:          uuid.New(),
		Name:        data["name"],
		Description: data["description"],
	}
	tx := db.DB.Create(&permission)
	if tx.Error != nil {
		log.Println(tx.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Error Creating Permission"})
	}

	return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Permission Created Successfully"})
}

func GetAllPermissions(c *fiber.Ctx) error {
	var permissions []m.Permission
	result := db.DB.Find(&permissions)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}

	return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Permissions Fetched Successfully", Data: permissions})
}
