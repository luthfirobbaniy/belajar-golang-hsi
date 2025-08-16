package middleware

import (
	"tugas-pertemuan-6/models"
	"tugas-pertemuan-6/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Jwt(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(401).JSON(models.ErrorResponse{
			Success: false,
			Message: "Authorization header is required",
		})
	}

	var tokenString string

	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		tokenString = authHeader[7:]
	}

	token, err := jwt.ParseWithClaims(
		tokenString,
		&utils.Claims{},
		func(t *jwt.Token) (any, error) {
			return utils.JwtSecret, nil
		},
	)

	if err != nil || !token.Valid {
		return c.Status(401).JSON(models.ErrorResponse{
			Success: false,
			Message: "Invalid or expired token",
		})
	}

	claims, ok := token.Claims.(*utils.Claims)

	if !ok {
		return c.Status(401).JSON(models.ErrorResponse{
			Success: false,
			Message: "Invalid token claims",
		})
	}

	c.Locals("id", claims.Id)
	c.Locals("username", claims.Username)
	c.Locals("role", claims.Role)

	return c.Next()
}
