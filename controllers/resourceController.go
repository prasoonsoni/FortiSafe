package controllers

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/db"
	m "github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

func hasPermission(permission_type string, user_id uuid.UUID, resource_id uuid.UUID) int {
	var permission m.Permission
	tx := db.DB.Where(&m.Permission{Name: permission_type}).Find(&permission)
	if tx.Error != nil {
		log.Println(tx.Error.Error())
		return 500
	}
	var user m.User
	tx = db.DB.Where(&m.User{ID: user_id}).Find(&user)
	if tx.Error != nil {
		log.Println(tx.Error.Error())
		return 500
	}
	var role m.Role
	tx = db.DB.Where(&m.Role{ID: user.RoleID}).Find(&role)
	if tx.Error != nil {
		log.Println(tx.Error.Error())
		return 500
	}
	if resource_id != uuid.Nil {
		var resource m.Resource
		db.DB.Where(&m.Resource{ID: resource_id}).Find(&resource)
		if !slices.Contains(resource.AssociatedRoles, user.RoleID.String()) {
			return 401
		}
	}
	if user.RoleID == uuid.Nil {
		return 401
	}

	if !slices.Contains(role.Permissions, permission.ID.String()) {
		return 401
	}
	return 200
}

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
	check := hasPermission("create", id, uuid.Nil)
	if check == 500 {
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	if check == 401 {
		return c.Status(fiber.StatusUnauthorized).JSON(&m.Response{Success: false, Message: "You don't have access to create resource"})
	}
	// Create Resource Here
	tx := db.DB.Create(&m.Resource{ID: uuid.New(), Name: resource.Name, Description: resource.Description, CreatedBy: id, AssociatedRoles: []string{c.Locals("user_role").(string)}})
	if tx.Error != nil {
		log.Println(tx.Error.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Resource Created Successfully"})
}

func GetResource(c *fiber.Ctx) error {
	// Get the user_id from the local context and cast it to a string
	user_id := c.Locals("user_id").(string)

	// Parse the user_id into a UUID
	id, err := uuid.Parse(user_id)

	// If error occurs parsing the used_id return Internal Server Error
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	resource_id, err := uuid.Parse(c.Params("resource_id"))
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	var resource m.Resource
	tx := db.DB.Where(&m.Resource{ID: resource_id}).Find(&resource)
	if tx.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	if tx.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(&m.Response{Success: false, Message: "Resource Not Found"})
	}

	check := hasPermission("read", id, resource_id)
	if check == 500 {
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	if check == 401 {
		return c.Status(fiber.StatusUnauthorized).JSON(&m.Response{Success: false, Message: "You don't have access to view resource"})
	}

	return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Resource Found Successfully", Data: resource})
}

func UpdateResource(c *fiber.Ctx) error {
	// Get the user_id from the local context and cast it to a string
	user_id := c.Locals("user_id").(string)

	// Parse the user_id into a UUID
	id, err := uuid.Parse(user_id)

	// If error occurs parsing the used_id return Internal Server Error
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	resource_id, err := uuid.Parse(c.Params("resource_id"))
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	var resource m.Resource
	tx := db.DB.Where(&m.Resource{ID: resource_id}).Find(&resource)
	if tx.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	if tx.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(&m.Response{Success: false, Message: "Resource Not Found"})
	}

	check := hasPermission("update", id, resource_id)
	if check == 500 {
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	if check == 401 {
		return c.Status(fiber.StatusUnauthorized).JSON(&m.Response{Success: false, Message: "You don't have access to update resource"})
	}

	var body m.Resource
	err = c.BodyParser(&body)
	if err != nil {
		// If there's an error in parsing the body, log the error and return an Internal Server Error response
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}

	resource.Name = body.Name
	resource.Description = body.Description

	tx = db.DB.Save(&resource)
	if tx.Error != nil {
		log.Println(tx.Error.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Error Updating Resource"})
	}
	return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Resource Deleted Successfully"})
}

func DeleteResource(c *fiber.Ctx) error {
	// Get the user_id from the local context and cast it to a string
	user_id := c.Locals("user_id").(string)

	// Parse the user_id into a UUID
	id, err := uuid.Parse(user_id)

	// If error occurs parsing the used_id return Internal Server Error
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	resource_id, err := uuid.Parse(c.Params("resource_id"))
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	var resource m.Resource
	tx := db.DB.Where(&m.Resource{ID: resource_id}).Find(&resource)
	if tx.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	if tx.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(&m.Response{Success: false, Message: "Resource Not Found"})
	}

	check := hasPermission("delete", id, resource_id)
	if check == 500 {
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	if check == 401 {
		return c.Status(fiber.StatusUnauthorized).JSON(&m.Response{Success: false, Message: "You don't have access to delete resource"})
	}

	tx = db.DB.Where(&m.Resource{ID: resource_id}).Delete(&m.Resource{})
	if tx.Error != nil {
		log.Println(tx.Error.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Error Deleting Resource"})
	}
	return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Resource Deleted Successfully"})
}

func AddAssociatedRoles(c *fiber.Ctx) error {
	var body m.AddAssociatedRolesBody
	err := c.BodyParser(&body)
	if err != nil {
		// If there's an error in parsing the body, log the error and return an Internal Server Error response
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	resource_id, err := uuid.Parse(body.ResourceID)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(&m.Response{Success: false, Message: "Give valid resource_id"})
	}
	var resource m.Resource
	tx := db.DB.Where(&m.Resource{ID: resource_id}).Find(&resource)
	if tx.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(&m.Response{Success: false, Message: "Resource Not Found"})
	}
	for _, role := range body.Roles {
		if !slices.Contains(resource.AssociatedRoles, role) {
			resource.AssociatedRoles = append(resource.AssociatedRoles, role)
		}
	}
	tx = db.DB.Save(&resource)
	if tx.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Role Added Successfully"})
}

func RemoveAssociatedRole(c *fiber.Ctx) error {
	var body m.RemoveAssociatedRoleBody
	err := c.BodyParser(&body)
	if err != nil {
		// If there's an error in parsing the body, log the error and return an Internal Server Error response
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	resource_id, err := uuid.Parse(body.ResourceID)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(&m.Response{Success: false, Message: "Give valid resource_id"})
	}
	var resource m.Resource
	tx := db.DB.Where(&m.Resource{ID: resource_id}).Find(&resource)
	if tx.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(&m.Response{Success: false, Message: "No resource Found"})
	}
	for i, role := range resource.AssociatedRoles {
		if role == body.RoleID {
			resource.AssociatedRoles = append(resource.AssociatedRoles[:i], resource.AssociatedRoles[i+1:]...)
		}
	}
	tx = db.DB.Save(&resource)
	if tx.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Error Removing Role"})
	}
	return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Role Removed Successfully"})

}

func BulkCreateResource(c *fiber.Ctx) error {
	// Get the user_id from the local context and cast it to a string
	user_id := c.Locals("user_id").(string)

	// Parse the user_id into a UUID
	id, err := uuid.Parse(user_id)

	// If error occurs parsing the used_id return Internal Server Error
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	check := hasPermission("create", id, uuid.Nil)
	if check == 500 {
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	if check == 401 {
		return c.Status(fiber.StatusUnauthorized).JSON(&m.Response{Success: false, Message: "You don't have access to create resource"})
	}
	file, err := c.FormFile("resources")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&m.Response{Success: false, Message: "Error Uploading File"})
	}
	src, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Failed to open the uploaded file"})
	}
	defer src.Close()
	// Create a CSV reader
	reader := csv.NewReader(src)
	// Process the CSV data
	var resources []*m.Resource
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Failed to read CSV file"})
		}
		if record[0] != "name" {
			name := record[0]
			description := record[1]
			associated_roles := strings.Split(strings.Trim(record[2], "{}"), ", ")
			created_by, _ := uuid.Parse(user_id)
			associated_roles = append(associated_roles, c.Locals("user_role").(string))
			fmt.Println(associated_roles)
			resource := m.Resource{
				ID:              uuid.New(),
				Name:            name,
				Description:     description,
				CreatedBy:       created_by,
				AssociatedRoles: associated_roles,
			}
			resources = append(resources, &resource)
		}
	}
	tx := db.DB.Create(&resources)
	if tx.Error != nil {
		// If there's an error in saving the user, log the error and return an Internal Server Error response
		log.Println(tx.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Error Creating Resources"})
	}
	return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Resources Created Successfully"})

}
