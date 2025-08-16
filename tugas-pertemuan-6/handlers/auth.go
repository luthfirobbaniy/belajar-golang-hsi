package handlers

import (
	"time"
	"tugas-pertemuan-6/models"
	"tugas-pertemuan-6/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	_ "tugas-pertemuan-6/docs"
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
		return c.Status(400).JSON(models.ErrorResponse{
			Success: false,
			Message: "Invalid request body!",
		})
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
		return c.Status(401).JSON(models.ErrorResponse{
			Success: false,
			Message: "Invalid credentials!",
		})
	}

	claims := utils.Claims{
		Id:       user.ID,
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(utils.JwtSecret)

	if err != nil {
		return c.Status(500).JSON(models.ErrorResponse{
			Success: false,
			Message: "Failed to generate token!",
		})
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
		return c.Status(400).JSON(models.ErrorResponse{
			Success: false,
			Message: "Invalid request body!",
		})
	}

	// Validate username with registered usernames
	for _, u := range users {
		if body.Username == u.Username {
			return c.Status(401).JSON(models.ErrorResponse{
				Success: false,
				Message: "Username has already exist!",
			})
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

	claims := utils.Claims{
		Id:       newUser.ID,
		Username: newUser.Username,
		Role:     newUser.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(utils.JwtSecret)

	if err != nil {
		return c.Status(500).JSON(models.ErrorResponse{
			Success: false,
			Message: "Failed to generate token!",
		})
	}

	return c.JSON(models.RegisterResponse{
		Success: true,
		Message: "Register successful!",
		Data: models.LoginData{
			Token: tokenString,
		},
	})
}
