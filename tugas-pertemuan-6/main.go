package main

import (
	"tugas-pertemuan-6/handlers"
	"tugas-pertemuan-6/middleware"
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
	Success bool      `json:"success" example:"true"`
	Message string    `json:"message" example:"Login successful!"`
	Data    LoginData `json:"data"`
}

// @Description Login response data
type LoginData struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Imx1dGhmaSIsImV4cCI6MTc1NTMzMTU5NiwiaWF0IjoxNzU1MjQ1MTk2fQ.7WktpMm0AyyfXUR5x68Om7Pps9uR1resDlh2bz9C_J8"`
}

// @Description Register response data
type RegisterRequest struct {
	Username string `json:"username" example:"luthfi"`
	Password string `json:"password" example:"123"`
}

// @Description Register response data
type RegisterResponse struct {
	Success bool      `json:"success" example:"true"`
	Message string    `json:"message" example:"Login successful!"`
	Data    LoginData `json:"data"`
}

// @Description Get students response
type GetStudentsResponse struct {
	Success bool             `json:"success" example:"true"`
	Message string           `json:"message" example:"Get students successful!"`
	Data    []models.Student `json:"data"`
}

// @Description Get student response
type GetStudentResponse struct {
	Success bool           `json:"success" example:"true"`
	Message string         `json:"message" example:"Get student successful!"`
	Data    models.Student `json:"data"`
}

// @Description Create student request
type CreateStudentRequest struct {
	NIM      string `json:"nim" example:"2021003"`
	Name     string `json:"name" example:"Budi"`
	Email    string `json:"email" example:"budi@example.com"`
	Major    string `json:"major" example:"Teknik Industri"`
	Semester int    `json:"semester" example:"2"`
}

// @Description Create student response
type CreateStudentResponse struct {
	Success bool           `json:"success" example:"true"`
	Message string         `json:"message" example:"Create student success!"`
	Data    models.Student `json:"data"`
}

// @Description Update student request
type UpdateStudentRequest struct {
	NIM      string `json:"nim" example:"2021001"`
	Name     string `json:"name" example:"Luthfi Edited"`
	Email    string `json:"email" example:"luthfi@example.com"`
	Major    string `json:"major" example:"Computer Science"`
	Semester int    `json:"semester" example:"1"`
}

// @Description Update student response
type UpdateStudentResponse struct {
	Success bool           `json:"success" example:"true"`
	Message string         `json:"message" example:"Update student success!"`
	Data    models.Student `json:"data"`
}

// @Description Delete student response
type DeleteStudentResponse struct {
	Success bool           `json:"success" example:"true"`
	Message string         `json:"message" example:"Delete student success!"`
	Data    models.Student `json:"data"`
}

// @Description Get profile response
type GetProfileResponse struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message" example:"Get profile success!"`
	Data    ProfileData `json:"data"`
}

// @Description Get profile response data
type ProfileData struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

// @Description Error response
type ErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Invalid request body"`
}

type Claims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func main() {
	app := fiber.New(fiber.Config{
		// ErrorHandler: func() {},
	})

	app.Get("/swagger/*", swagger.HandlerDefault)

	// Authentication Endpoints
	app.Post("/api/auth/login", handlers.Login)
	app.Post("/api/auth/register", handlers.Register)

	// Student Management Endpoints (Protected)
	app.Get("/api/students", middleware.Jwt, handlers.GetStudents)
	app.Get("/api/students/:id", middleware.Jwt, handlers.GetStudent)
	app.Post("/api/students", middleware.Jwt, handlers.CreateStudent)
	app.Put("/api/students/:id", middleware.Jwt, handlers.UpdateStudent)
	app.Delete("/api/students/:id", middleware.Jwt, handlers.DeleteStudent)

	// Profile Endpoint
	app.Get("/api/profile", middleware.Jwt, GetProfile)

	app.Listen(":3000")
}

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

// GetProfile godoc
// @Summary Get profile
// @Description Get profile
// @Tags Profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} GetProfileResponse "Get profile successful"
// @Failure 401 {object} ErrorResponse "Invalid credentials"
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

	return c.JSON(GetProfileResponse{
		Success: true,
		Message: "Get profile success!",
		Data: ProfileData{
			ID:       user.ID,
			Username: user.Username,
			Role:     user.Role,
		},
	})
}
