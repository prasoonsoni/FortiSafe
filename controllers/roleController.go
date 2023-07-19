package controllers

import (
	"log"

	"github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/db"
	m "github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CreateRoleBody struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Permissions []string `json:"permissions"`
}

func CreateRole(c *fiber.Ctx) error {
	var data CreateRoleBody
	// Parse the request body into the data variable
	err := c.BodyParser(&data)
	if err != nil {
		// If there's an error in parsing the body, log the error and return an Internal Server Error response
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	var role *m.Role
	result := db.DB.Where(&m.Permission{Name: data.Name}).Find(&role)
	if result.RowsAffected == 1 {
		return c.Status(fiber.StatusOK).JSON(&m.Response{Success: false, Message: "Role already exists"})
	}

	var permission_ids []uuid.UUID
	for _, permission := range data.Permissions {
		id, _ := uuid.Parse(string(permission))
		permission_ids = append(permission_ids, id)
	}

	var permissions []m.Permission
	for _, permission_id := range permission_ids {
		tmp_permission := m.Permission{
			ID: permission_id,
		}
		permissions = append(permissions, tmp_permission)
	}
	role = &m.Role{
		ID:          uuid.New(),
		Name:        data.Name,
		Description: data.Description,
		Permissions: permissions,
	}
	tx := db.DB.Create(&role)
	if tx.Error != nil {
		log.Println(tx.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Error Creating Role"})
	}

	return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Role Created Successfully"})
}


