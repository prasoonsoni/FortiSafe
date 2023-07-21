package controllers

import (
	"log"

	"github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/db"
	m "github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateGroup(c *fiber.Ctx) error {
	var data m.CreateGroupBody
	// Parse the request body into the data variable
	err := c.BodyParser(&data)
	log.Println(data)
	if err != nil {
		// If there's an error in parsing the body, log the error and return an Internal Server Error response
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}

	group_id := uuid.New()
	tx := db.DB.Create(&m.Group{
		ID:          group_id,
		Name:        data.Name,
		Description: data.Description,
		Permissions: data.Permissions,
	})
	if tx.Error != nil {
		log.Println(tx.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Error Creating Group"})
	}

	return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Group Created Successfully"})
}
