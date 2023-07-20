package controllers

import (
	"fmt"
	"log"

	"github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/db"
	m "github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type CreateRoleBody struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Permissions []string `json:"permissions"`
}

func CreateRole(c *fiber.Ctx) error {
	var data CreateRoleBody
	// Parse the request body into the data variable
	err := c.BodyParser(&data)
	log.Println(data)
	if err != nil {
		// If there's an error in parsing the body, log the error and return an Internal Server Error response
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	var role *m.Role
	result := db.DB.Where(&m.Role{Name: data.Name}).Find(&role)
	if result.RowsAffected == 1 {
		return c.Status(fiber.StatusOK).JSON(&m.Response{Success: false, Message: "Role already exists"})
	}

	role_id := uuid.New()
	tx := db.DB.Create(&m.Role{
		ID:          role_id,
		Name:        data.Name,
		Description: data.Description,
		Permissions: data.Permissions,
	})
	if tx.Error != nil {
		log.Println(tx.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Error Creating Role"})
	}

	for _, permission := range data.Permissions {
		permission_id, _ := uuid.Parse(permission)
		_ = db.DB.Create(&m.RolePermission{
			ID:           uuid.New(),
			RoleID:       role_id,
			PermissionID: permission_id,
		})
	}

	return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Role Created Successfully"})
}

type AddPermissionBody struct {
	RoleID      string   `json:"role_id"`
	Permissions []string `json:"permissions"`
}

func AddPermission(c *fiber.Ctx) error {
	var data AddPermissionBody
	// Parse the request body into the data variable
	err := c.BodyParser(&data)
	log.Println(data)
	if err != nil {
		// If there's an error in parsing the body, log the error and return an Internal Server Error response
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	role_id, err := uuid.Parse(data.RoleID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Give Valid ID"})
	}
	var role *m.Role
	tx := db.DB.Where(&m.Role{ID: role_id}).Find(&role)
	if tx.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(&m.Response{Success: false, Message: "Role Not Found"})
	}

	for _, permission := range data.Permissions {
		if !slices.Contains(role.Permissions, permission) {
			id, _ := uuid.Parse(permission)
			_ = db.DB.Create(&m.RolePermission{
				ID:           uuid.New(),
				RoleID:       role_id,
				PermissionID: id,
			})
		}
	}

	for _, permission := range data.Permissions {
		if !slices.Contains(role.Permissions, permission) {
			role.Permissions = append(role.Permissions, permission)
		}
	}
	tx = db.DB.Save(&role)
	if tx.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Error Adding Permission"})

	}
	return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Permissions Added Successfully"})
}

func GetAllRoles(c *fiber.Ctx) error {
	var roles []m.Role
	result := db.DB.Find(&roles)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	var data []interface{}
	for _, role := range roles {
		var tmp_data = make(map[string]interface{})
		fmt.Println(role)
		tmp_data["ID"] = role.ID
		tmp_data["created_at"] = role.CreatedAt
		tmp_data["updated_at"] = role.UpdatedAt
		tmp_data["name"] = role.Name
		tmp_data["description"] = role.Description
		var permissions []m.Permission
		for _, permission := range role.Permissions {
			id, _ := uuid.Parse(permission)
			var tmp_permission m.Permission
			_ = db.DB.Where(&m.Permission{ID: id}).Find(&tmp_permission)
			permissions = append(permissions, tmp_permission)
		}
		tmp_data["permissions"] = permissions
		data = append(data, tmp_data)
	}
	return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Permissions Fetched Successfully", Data: data})
}
