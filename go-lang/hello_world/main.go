package main

import (
    "github.com/gofiber/fiber/v2"
)

func main() {
    // Create a new Fiber app
    app := fiber.New()

    // Define a route for the root path
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    // Start the server on port 3000
    app.Listen(":3000")
}