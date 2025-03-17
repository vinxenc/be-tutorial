package users

import (
	"github.com/Oudwins/zog"
	"github.com/gofiber/fiber/v2"
)

func ValidateCreateUserRequest() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := new(User)

		c.BodyParser(user)

		if issues := user.Validate(); len(issues) > 0 {
			errors := []zog.ZogIssue{}
			for name, issue := range issues {
				if name == "$first" {
					continue
				}

				errors = append(errors, *issue[0])
			}

			return c.Status(400).JSON(fiber.Map{
				"errors": errors,
			})
		}

		return c.Next()
	}
}
