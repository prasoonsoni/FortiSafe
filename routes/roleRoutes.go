package routes

import (
	roleControllers "github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/controllers"
	"github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupRoleRoutes(app *fiber.App) {
	app.Post("/api/role/create", middlewares.AuthenticateAdmin, roleControllers.CreateRole)
	app.Put("/api/role/permission/add", middlewares.AuthenticateAdmin, roleControllers.AddPermission)
	app.Delete("/api/role/permission/remove", middlewares.AuthenticateAdmin, roleControllers.RemovePermission)
	app.Get("/api/role/get/all", middlewares.AuthenticateAdmin, roleControllers.GetAllRoles)
	app.Put("/api/role/assign", middlewares.AuthenticateAdmin, roleControllers.AssignRole)
	app.Put("/api/role/unassign", middlewares.AuthenticateAdmin, roleControllers.UnassignRole)
}
