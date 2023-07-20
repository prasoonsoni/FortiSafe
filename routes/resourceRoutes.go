package routes

import (
	resourceControlers "github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/controllers"
	"github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupResourceRoutes(app *fiber.App) {
	app.Post("/api/resource/create", middlewares.AuthenticateUser, resourceControlers.CreateResource)
}
