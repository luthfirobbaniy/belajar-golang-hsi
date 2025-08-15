package main

import (
	"strconv"
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
	Status  int       `json:"status" example:"200"`
	Message string    `json:"message" example:"Login successful!"`
	Data    LoginData `json:"data"`
}

// @Description Login response data
type LoginData struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Imx1dGhmaSIsImV4cCI6MTc1NTMzMTU5NiwiaWF0IjoxNzU1MjQ1MTk2fQ.7WktpMm0AyyfXUR5x68Om7Pps9uR1resDlh2bz9C_J8"`
}

// @Description Get students response
type GetStudentsResponse struct {
	Status  int              `json:"status" example:"200"`
	Message string           `json:"message" example:"Get students successful!"`
	Data    []models.Student `json:"data"`
}

// @Description Get student response
type GetStudentResponse struct {
	Status  int            `json:"status" example:"200"`
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

// @Description Error response
type ErrorResponse struct {
	Success bool   `json:"success" example:"false"`
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

	// Student Management Endpoints (Protected)
	app.Get("/api/students", jwtMiddleware, getStudents)
	app.Get("/api/students/:id", jwtMiddleware, getStudent)
	app.Post("/api/students", jwtMiddleware, createStudent)

	app.Listen(":3000")
}

var jwtSecret = []byte("jwt-secret")

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
	{
		ID:       3,
		Username: "luthfi",
		Password: "123",
		Role:     "student",
	},
}

var students = []models.Student{
	{
		ID:       1,
		NIM:      "2021001",
		Name:     "Luthfi",
		Email:    "luthfi@example.com",
		Major:    "Computer Science",
		Semester: 1,
	},
	{
		ID:       2,
		NIM:      "2021002",
		Name:     "Ahmad",
		Email:    "ahmad@example.com",
		Major:    "Electrical Engineering",
		Semester: 1,
	},
}

var latestId = 2

func jwtMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "Authorization header is required",
		})
	}

	var tokenString string

	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		tokenString = authHeader[7:]
	}

	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(t *jwt.Token) (any, error) {
			return jwtSecret, nil
		},
	)

	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})
	}

	claims, ok := token.Claims.(*Claims)

	if !ok {
		return c.Status(401).JSON(fiber.Map{
			"message": "Invalid token claims",
		})
	}

	c.Locals("username", claims.Username)

	return c.Next()
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

	var user models.User

	// Find user
	for _, u := range users {
		if req.Username == u.Username && req.Password == u.Password {
			user = u
			break
		}
	}

	if user.ID == 0 {
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
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "Failed to generate token",
		})
	}

	return c.JSON(fiber.Map{
		"status":  200,
		"message": "Login successful!",
		"data": fiber.Map{
			"token": tokenString,
		},
	})
}

// getStudents godoc
// @Summary Get all student
// @Description Get all student data
// @Tags Students
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} GetStudentsResponse "Get students successful"
// @Failure 401 {object} ErrorResponse "Authorization header required or Invalid token"
// @Router /students [get]
func getStudents(c *fiber.Ctx) error {
	return c.JSON(GetStudentsResponse{
		Status:  200,
		Message: "Get students successful!",
		Data:    students,
	})
}

// getStudent godoc
// @Summary Get one student
// @Description Get one student data
// @Tags Students
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Student id"
// @Success 200 {object} GetStudentResponse "Get student successful"
// @Failure 400 {object} ErrorResponse "Missing parameter"
// @Failure 401 {object} ErrorResponse "Authorization header required or Invalid token"
// @Failure 404 {object} ErrorResponse "Student not found"
// @Router /students/{id} [get]
func getStudent(c *fiber.Ctx) error {
	if c.Params("id") == "" {
		return c.Status(400).JSON(ErrorResponse{
			Status:  400,
			Message: "Missing parameter!",
		})
	}

	var student models.Student
	// student := &models.Student{}

	for _, s := range students {
		if id, _ := strconv.Atoi(c.Params("id")); s.ID == id {
			student = s
			break
		}
	}

	if student.ID == 0 {
		return c.Status(404).JSON(ErrorResponse{
			Status:  404,
			Message: "Student not found!",
		})
	}

	return c.JSON(GetStudentResponse{
		Status:  200,
		Message: "Get student successful!",
		Data:    student,
	})
}

// createStudent godoc
// @Summary Create student
// @Description Create student
// @Tags Students
// @Accept json
// @Produce json
// @Param request body CreateStudentRequest true "Create student data"
// @Success 200 {object} CreateStudentResponse "Create student successful"
// @Failure 400 {object} ErrorResponse "Invalid request body"
// @Failure 401 {object} ErrorResponse "Invalid credentials"
// @Router /students [post]
func createStudent(c *fiber.Ctx) error {
	var body CreateStudentRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(ErrorResponse{
			Success: false,
			Message: "",
		})
	}

	latestId += 1

	newStudent := models.Student{
		ID:       latestId,
		NIM:      body.NIM,
		Name:     body.Name,
		Email:    body.Email,
		Major:    body.Major,
		Semester: body.Semester,
	}

	students = append(students, newStudent)

	return c.JSON(CreateStudentResponse{
		Success: true,
		Message: "Create student success!",
		Data:    newStudent,
	})
}
