package users

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserController struct{}

func NewController() *UserController {
	return &UserController{}
}

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
	user := new(User)
	ctx.BodyParser(user)

	// Convert int ID to string for map key
	idStr := strconv.Itoa(user.ID)
	users[idStr] = *user
	return ctx.Status(201).JSON(user)
}

func (c *UserController) GetUsers(ctx *fiber.Ctx) error {
	userList := make([]User, 0, len(users))
	for _, user := range users {
		userList = append(userList, user)
	}
	return ctx.JSON(userList)
}

func (c *UserController) GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, exists := users[id]
	if !exists {
		return ctx.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	return ctx.JSON(user)
}

func (c *UserController) UpdateUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if _, exists := users[id]; !exists {
		return ctx.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	user := new(User)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "Invalid user data",
		})
	}

	users[id] = *user
	return ctx.JSON(user)
}

func (c *UserController) DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if _, exists := users[id]; !exists {
		return ctx.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	delete(users, id)
	return ctx.SendStatus(204)
}
