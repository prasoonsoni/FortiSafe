package routes

import (
	resourceControllers "github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/controllers"
	"github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupResourceRoutes(app *fiber.App) {
	app.Post("/api/resource/create", middlewares.AuthenticateUser, resourceControllers.CreateResource)
	app.Get("/api/resource/get/:resource_id", middlewares.AuthenticateUser, resourceControllers.GetResource)
	app.Put("/api/resource/update/:resource_id", middlewares.AuthenticateUser, resourceControllers.UpdateResource)
	app.Delete("/api/resource/delete/:resource_id", middlewares.AuthenticateUser, resourceControllers.DeleteResource)
}
