package main

import (
	"day-003-middleware/src/middlewares"
	"day-003-middleware/src/modules/users"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Add logger middleware for all routes
	app.Use(middlewares.LoggerMiddleware())

	// API routes group
	api := app.Group("/api")

	// Setup routes using the users package
	users.RegisterRoutes(api)

	log.Fatal(app.Listen(":3000"))
}
