package main

import (
	"github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/db"
	m "github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/models"
	"github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-prasoonsoni/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Initialize a new Fiber application
	app := fiber.New()

	// Attach a logger middleware to the app to log every incoming HTTP request
	app.Use(logger.New())

	// Set up Cross-Origin Resource Sharing (CORS) for the app
	// It's allowing every request from all origins that require credentials
	app.Use(cors.New(cors.Config{AllowOrigins: "*", AllowCredentials: true}))
	// Establishing Connection to Database
	db.Connect()
	// Set up a route for the root ("/") path
	// When a GET request is received at this path, the server sends a JSON response indicating that it's working
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(m.Response{Success: true, Message: "Server is Working"})
	})
	routes.SetupUserRoutes(app)
	routes.SetupPermissionRoutes(app)
	routes.SetupRoleRoutes(app)
	// Start the server and make it listen for incoming HTTP requests on port 3000
	app.Listen(":8080")

}
