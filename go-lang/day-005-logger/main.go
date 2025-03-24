package main

import (
	log "day-005-logger/src/libs/logger"
	"day-005-logger/src/middlewares"
	"day-005-logger/src/modules/users"

	"github.com/gofiber/fiber/v2"
)

func main() {
	logger := log.NewLogger(&log.LoggerOptions{
		Transports: []log.Transport{log.NewConsoleTransport()},
	})

	app := fiber.New()

	// Add logger middleware for all routes
	app.Use(middlewares.LoggerMiddleware())

	// API routes group
	api := app.Group("/api")

	// Setup routes using the users package
	users.RegisterRoutes(api)

	logger.Info("Server started on port 3000")

	app.Listen(":3000")
}
