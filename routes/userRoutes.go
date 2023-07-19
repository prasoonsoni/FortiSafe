package userRoutes

import (
	userControllers "github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/controllers"
	"github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/user/create", userControllers.CreateUser)
	app.Post("/api/user/login", userControllers.LoginUser)
	app.Get("/api/user/get", middlewares.AuthenticateUser, userControllers.GetUser)
	app.Put("/api/user/deactivate", middlewares.AuthenticateUser, userControllers.DeactivateUser)
	app.Put("/api/user/activate", middlewares.AuthenticateUser, userControllers.ActivateUser)
}
