package routes

import (
	roleControllers "github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoleRoutes(app *fiber.App) {
	app.Post("/api/role/create", roleControllers.CreateRole)
	app.Put("/api/role/permission/add", roleControllers.AddPermission)
}
