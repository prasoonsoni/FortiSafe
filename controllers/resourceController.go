package controllers

import (
	"log"

	"github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/db"
	m "github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

func CreateResource(c *fiber.Ctx) error {
	var resource m.Resource
	err := c.BodyParser(&resource)
	if err != nil {
		// If there's an error in parsing the body, log the error and return an Internal Server Error response
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	// Get the user_id from the local context and cast it to a string
	user_id := c.Locals("user_id").(string)

	// Parse the user_id into a UUID
	id, err := uuid.Parse(user_id)

	// If error occurs parsing the used_id return Internal Server Error
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	var permission m.Permission
	tx := db.DB.Where(&m.Permission{Name: "create"}).Find(&permission)
	if tx.Error != nil {
		log.Println(tx.Error.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	var user m.User
	tx = db.DB.Where(&m.User{ID: id}).Find(&user)
	if tx.Error != nil {
		log.Println(tx.Error.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	var role m.Role
	tx = db.DB.Where(&m.Role{ID: user.RoleID}).Find(&role)
	if tx.Error != nil {
		log.Println(tx.Error.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	if !slices.Contains(role.Permissions, permission.ID.String()) {
		return c.Status(fiber.StatusUnauthorized).JSON(&m.Response{Success: false, Message: "You don't have access to create resource"})
	}
	// Create Resource Here
	tx = db.DB.Create(&m.Resource{ID: uuid.New(), Name: resource.Name, Description: resource.Description, CreatedBy: id, AssociatedRoles: []string{role.ID.String()}})
	if tx.Error != nil {
		log.Println(tx.Error.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Resource Created Successfully"})
}
