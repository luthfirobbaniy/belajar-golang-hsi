package main

import (
	"log"
	"time"
	"tugas-pertemuan-6/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/golang-jwt/jwt/v5"

	_ "tugas-pertemuan-6/docs"
)

// @title Sistem Manajemen Mahasiswa
// @version 1.0
// @description Sistem manajemen mahasiswa using Go Fiber
// @host localhost:3000
// @BasePath /api
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

// @Description Login request payload
type LoginRequest struct {
	Username string `json:"username" example:"luthfi"`
	Password string `json:"password" example:"123"`
}

// @Description Login response
type LoginResponse struct {
	Status  int       `json:"status" example:"1"`
	Message string    `json:"message" example:"Invalid request body"`
	Data    LoginData `json:"data"`
}

// @Description Login response data
type LoginData struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Imx1dGhmaSIsImV4cCI6MTc1NTMzMTU5NiwiaWF0IjoxNzU1MjQ1MTk2fQ.7WktpMm0AyyfXUR5x68Om7Pps9uR1resDlh2bz9C_J8"`
}

type ErrorResponse struct {
	Status  int    `json:"status" example:"1"`
	Message string `json:"message" example:"Invalid request body"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func main() {
	app := fiber.New(fiber.Config{
		// ErrorHandler: func() {},
	})

	app.Get("/swagger/*", swagger.HandlerDefault)

	// Authentication Endpoints
	app.Post("/api/auth/login", login)

	app.Listen(":3000")
}

var jwtSecret = []byte("jwt-secret")

var user = models.User{
	ID:       1,
	Username: "luthfi",
	Password: "123",
	Role:     "student",
}

// login godoc
// @Summary User login
// @Description Authenticate user with static credentials and return JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body LoginRequest true "Login credentials"
// @Success 200 {object} LoginResponse "Login successful"
// @Failure 400 {object} ErrorResponse "Invalid request body"
// @Failure 401 {object} ErrorResponse "Invalid credentials"
// @Failure 500 {object} ErrorResponse "Failed to generate token"
// @Router /auth/login [post]
func login(c *fiber.Ctx) error {
	var req LoginRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid request body",
		})
	}

	if req.Username != user.Username && req.Password != user.Password {
		return c.Status(401).JSON(fiber.Map{
			"status":  401,
			"message": "Invalid credentials",
		})
	}

	claims := Claims{
		Username: req.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "Failed to generate token",
		})
	}

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Login Successful!",
		"data": fiber.Map{
			"token": tokenString,
		},
	})
}
