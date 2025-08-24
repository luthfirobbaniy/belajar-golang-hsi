package handlers

import (
	"tugas-pertemuan-7/models"
	"tugas-pertemuan-7/utils"

	"github.com/gofiber/fiber/v2"

	_ "tugas-pertemuan-7/docs"
)

var users = []models.User{
	{
		ID:       1,
		Username: "admin",
		Password: "admin123",
		Role:     "admin",
	},
	{
		ID:       2,
		Username: "student1",
		Password: "student123",
		Role:     "student",
	},
}

var latestUserId = 2

// login godoc
// @Summary User login
// @Description Authenticate user with static credentials and return JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body models.LoginRequest true "Login credentials"
// @Success 200 {object} models.LoginResponse "Login successful"
// @Failure 400 {object} models.ErrorResponse "Invalid request body"
// @Failure 401 {object} models.ErrorResponse "Invalid credentials"
// @Failure 500 {object} models.ErrorResponse "Failed to generate token"
// @Router /auth/login [post]
func Login(c *fiber.Ctx) error {
	var body models.LoginRequest

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body!")
	}

	var user models.User

	// Find user
	for _, u := range users {
		if body.Username == u.Username && body.Password == u.Password {
			user = u
			break
		}
	}

	// If user not found
	if user.ID == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid credentials!")
	}

	tokenString, err := utils.CreateJwt(&user)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to generate token!")
	}

	return c.JSON(models.LoginResponse{
		Success: true,
		Message: "Login successful!",
		Data: models.LoginData{
			Token: tokenString,
		},
	})
}

// register godoc
// @Summary User register
// @Description Register new user and return JWT token for auto login
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body models.RegisterRequest true "Register credentials"
// @Success 200 {object} models.RegisterResponse "Register successful"
// @Failure 400 {object} models.ErrorResponse "Invalid request body"
// @Failure 401 {object} models.ErrorResponse "Invalid credentials"
// @Failure 500 {object} models.ErrorResponse "Failed to generate token"
// @Router /auth/register [post]
func Register(c *fiber.Ctx) error {
	var body models.RegisterRequest

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body!")
	}

	// Validate username with registered usernames
	for _, u := range users {
		if body.Username == u.Username {
			return fiber.NewError(fiber.StatusBadRequest, "Username has already exist!")
		}
	}

	latestUserId += 1

	newUser := models.User{
		ID:       latestUserId,
		Username: body.Username,
		Password: body.Password,
		Role:     "student", // Only accept "user" role for registration
	}

	// Store new user to db
	users = append(users, newUser)

	tokenString, err := utils.CreateJwt(&newUser)

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to generate token!")
	}

	return c.Status(201).JSON(models.RegisterResponse{
		Success: true,
		Message: "Register successful!",
		Data: models.LoginData{
			Token: tokenString,
		},
	})
}
