package handlers

import (
	"tugas-pertemuan-6/models"

	"github.com/gofiber/fiber/v2"

	_ "tugas-pertemuan-6/docs"
)

// GetProfile godoc
// @Summary Get profile
// @Description Get profile
// @Tags Profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.GetProfileResponse "Get profile successful"
// @Failure 401 {object} models.ErrorResponse "Invalid credentials"
// @Router /profile [get]
func GetProfile(c *fiber.Ctx) error {
	id := c.Locals("id")

	var user models.User

	// Find profile
	for _, u := range users {
		if u.ID == id {
			user = u
		}
	}

	return c.JSON(models.GetProfileResponse{
		Success: true,
		Message: "Get profile success!",
		Data: models.ProfileData{
			ID:       user.ID,
			Username: user.Username,
			Role:     user.Role,
		},
	})
}
