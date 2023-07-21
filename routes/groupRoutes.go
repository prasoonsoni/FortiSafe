package routes

import (
	groupControllers "github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/controllers"
	"github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupGroupRoutes(app *fiber.App) {
	app.Post("/api/group/create", middlewares.AuthenticateAdmin, groupControllers.CreateGroup)
	app.Put("/api/group/permission/add", middlewares.AuthenticateAdmin, groupControllers.AddGroupPermission)
}
