package middleware

import (
	"tugas-pertemuan-7/utils"

	"github.com/gofiber/fiber/v2"
)

func Jwt(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Authorization header is required")
	}

	claims, err := utils.ParseJwt(authHeader)

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	c.Locals("id", claims.Id)
	c.Locals("username", claims.Username)
	c.Locals("role", claims.Role)

	return c.Next()
}
