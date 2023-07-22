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

func CreateRole(c *fiber.Ctx) error {
	var data m.CreateRoleBody
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

func AddPermission(c *fiber.Ctx) error {
	var data m.AddPermissionBody
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

func RemovePermission(c *fiber.Ctx) error {
	var data m.DeletePermissionBody
	err := c.BodyParser(&data)
	log.Println(data)
	if err != nil {
		// If there's an error in parsing the body, log the error and return an Internal Server Error response
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	role_id, _ := uuid.Parse(data.RoleID)
	permission_id, _ := uuid.Parse(data.PermissionID)
	tx := db.DB.Where(&m.RolePermission{RoleID: role_id, PermissionID: permission_id}).Delete(&m.RolePermission{})
	if tx.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Error Removing Permission"})
	}
	var role m.Role
	_ = db.DB.Where(&m.Role{ID: role_id}).Find(&role)
	for i, permission := range role.Permissions {
		if permission == data.PermissionID {
			role.Permissions = append(role.Permissions[:i], role.Permissions[i+1:]...)
		}
	}
	tx = db.DB.Save(&role)
	if tx.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Error Removing Permission"})
	}
	return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Permission Removed Successfully"})
}

func AssignRole(c *fiber.Ctx) error {
	var data m.AssignRoleBody
	err := c.BodyParser(&data)
	log.Println(data)
	if err != nil {
		// If there's an error in parsing the body, log the error and return an Internal Server Error response
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	user_id, err := uuid.Parse(data.UserID)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Give valid user_id"})
	}
	if tx := db.DB.Where(m.User{ID: user_id}).Find(&m.User{}); tx.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(&m.Response{Success: false, Message: "User not found"})
	}
	role_id, err := uuid.Parse(data.RoleID)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Give valid role_id"})
	}
	if tx := db.DB.Where(m.Role{ID: role_id}).Find(&m.Role{}); tx.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(&m.Response{Success: false, Message: "Role not found"})
	}
	var user m.User
	tx := db.DB.Where(m.User{ID: user_id}).Find(&user)
	if tx.Error != nil {
		log.Println(tx.Error.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Error Assigning Role"})
	}
	user.RoleID = role_id
	tx = db.DB.Save(&user)
	if tx.Error != nil {
		log.Println(tx.Error.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Error Assigning Role"})
	}
	return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Role Assigned Successfully"})
}

func UnassignRole(c *fiber.Ctx) error {
	user_id, err := uuid.Parse(c.Query("user_id"))
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Give valid user_id"})
	}
	if tx := db.DB.Where(m.User{ID: user_id}).Find(&m.User{}); tx.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(&m.Response{Success: false, Message: "User not found"})
	}
	var user m.User
	tx := db.DB.Where(&m.User{ID: user_id}).Find(&user)
	if tx.Error != nil {
		log.Println(tx.Error.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Error Unassign Role"})
	}
	user.RoleID = uuid.Nil
	tx = db.DB.Save(&user)

	if tx.Error != nil {
		log.Println(tx.Error.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Error unassign Role"})
	}
	return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Role Unassigned Successfully"})
}
