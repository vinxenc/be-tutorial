package users

import (
    "github.com/gofiber/fiber/v2"
)

// RegisterRoutes registers all user routes
func RegisterRoutes(api fiber.Router) {
    userController := NewController()
    
    // User routes group
    userRoutes := api.Group("/users")
    
    userRoutes.Post("/", userController.CreateUser)
    userRoutes.Get("/", userController.GetUsers)
    userRoutes.Get("/:id", userController.GetUser)
    userRoutes.Put("/:id", userController.UpdateUser)
    userRoutes.Delete("/:id", userController.DeleteUser)
}