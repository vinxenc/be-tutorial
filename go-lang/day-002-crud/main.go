package main

import (
    "github.com/gofiber/fiber/v2"
    "log"
    "day-002-crud/src/modules/users"
)

func main() {
    app := fiber.New()
    
    // API routes group
    api := app.Group("/api")
    
    // Setup routes using the users package
    users.RegisterRoutes(api)

    log.Fatal(app.Listen(":3000"))
}