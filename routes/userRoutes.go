package userRoutes

import (
	userControllers "github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/user/create", userControllers.CreateUser)
	app.Post("/api/user/login", userControllers.LoginUser)
}