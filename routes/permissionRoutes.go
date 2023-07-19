package routes

import (
	permissionControllers "github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupPermissionRoutes(app *fiber.App) {
	app.Post("/api/permission/create", permissionControllers.CreatePermission)
	app.Get("/api/permission/all", permissionControllers.GetAllPermissions)
}
