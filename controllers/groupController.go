package controllers

import (
	"log"

	"github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/db"
	m "github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
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

func AddGroupPermission(c *fiber.Ctx) error {
	var data m.AddGroupPermissionBody
	// Parse the request body into the data variable
	err := c.BodyParser(&data)
	log.Println(data)
	if err != nil {
		// If there's an error in parsing the body, log the error and return an Internal Server Error response
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	group_id, err := uuid.Parse(data.GroupID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Give Valid ID"})
	}
	var group *m.Group
	tx := db.DB.Where(&m.Group{ID: group_id}).Find(&group)
	if tx.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(&m.Response{Success: false, Message: "Group Not Found"})
	}

	for _, permission := range data.Permissions {
		if !slices.Contains(group.Permissions, permission) {
			group.Permissions = append(group.Permissions, permission)
		}
	}
	tx = db.DB.Save(&group)
	if tx.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Error Adding Permission"})

	}
	return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Permissions Added Successfully"})
}

func RemoveGroupPermission(c *fiber.Ctx) error {
	var data m.DeleteGroupPermissionBody
	err := c.BodyParser(&data)
	log.Println(data)
	if err != nil {
		// If there's an error in parsing the body, log the error and return an Internal Server Error response
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Internal Server Error"})
	}
	group_id, _ := uuid.Parse(data.GroupID)
	var group m.Group
	_ = db.DB.Where(&m.Group{ID: group_id}).Find(&group)
	for i, permission := range group.Permissions {
		if permission == data.PermissionID {
			group.Permissions = append(group.Permissions[:i], group.Permissions[i+1:]...)
		}
	}
	tx := db.DB.Save(&group)
	if tx.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&m.Response{Success: false, Message: "Error Removing Permission"})
	}
	return c.Status(fiber.StatusOK).JSON(&m.Response{Success: true, Message: "Permission Removed Successfully"})
}
