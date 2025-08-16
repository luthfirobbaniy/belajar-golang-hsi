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
	app.Post("/api/auth/login", login)
	app.Post("/api/auth/register", register)

	// Student Management Endpoints (Protected)
	app.Get("/api/students", jwtMiddleware, getStudents)
	app.Get("/api/students/:id", jwtMiddleware, getStudent)
	app.Post("/api/students", jwtMiddleware, createStudent)
	app.Put("/api/students/:id", jwtMiddleware, updateStudent)
	app.Delete("/api/students/:id", jwtMiddleware, DeleteStudent)

	// Profile Endpoint
	app.Get("/api/profile", jwtMiddleware, GetProfile)

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

var latestUserId = 2
var latestStudentId = 2

func jwtMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(401).JSON(ErrorResponse{
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
		&Claims{},
		func(t *jwt.Token) (any, error) {
			return jwtSecret, nil
		},
	)

	if err != nil || !token.Valid {
		return c.Status(401).JSON(ErrorResponse{
			Success: false,
			Message: "Invalid or expired token",
		})
	}

	claims, ok := token.Claims.(*Claims)

	if !ok {
		return c.Status(401).JSON(ErrorResponse{
			Success: false,
			Message: "Invalid token claims",
		})
	}

	c.Locals("id", claims.Id)
	c.Locals("username", claims.Username)
	c.Locals("role", claims.Role)

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
		return c.Status(400).JSON(ErrorResponse{
			Success: false,
			Message: "Invalid request body!",
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
		return c.Status(401).JSON(ErrorResponse{
			Success: false,
			Message: "Invalid credentials!",
		})
	}

	claims := Claims{
		Id:       user.ID,
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		return c.Status(500).JSON(ErrorResponse{
			Success: false,
			Message: "Failed to generate token!",
		})
	}

	return c.JSON(LoginResponse{
		Success: true,
		Message: "Login successful!",
		Data: LoginData{
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
// @Param credentials body RegisterRequest true "Register credentials"
// @Success 200 {object} RegisterResponse "Register successful"
// @Failure 400 {object} ErrorResponse "Invalid request body"
// @Failure 401 {object} ErrorResponse "Invalid credentials"
// @Failure 500 {object} ErrorResponse "Failed to generate token"
// @Router /auth/register [post]
func register(c *fiber.Ctx) error {
	var body RegisterRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(ErrorResponse{
			Success: false,
			Message: "Invalid request body!",
		})
	}

	// Validate username with registered usernames
	for _, u := range users {
		if body.Username == u.Username {
			return c.Status(401).JSON(ErrorResponse{
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

	claims := Claims{
		Id:       newUser.ID,
		Username: newUser.Username,
		Role:     newUser.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		return c.Status(500).JSON(ErrorResponse{
			Success: false,
			Message: "Failed to generate token!",
		})
	}

	return c.JSON(RegisterResponse{
		Success: true,
		Message: "Register successful!",
		Data: LoginData{
			Token: tokenString,
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
		Success: true,
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
			Success: false,
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
			Success: false,
			Message: "Student not found!",
		})
	}

	return c.JSON(GetStudentResponse{
		Success: true,
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
// @Security BearerAuth
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
			Message: "Body parsing failed!",
		})
	}

	latestUserId += 1

	newStudent := models.Student{
		ID:       latestStudentId,
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

// updateStudent godoc
// @Summary Update student
// @Description Update student
// @Tags Students
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body UpdateStudentRequest true "Update student data"
// @Param id path int true "Student id"
// @Success 200 {object} UpdateStudentResponse "Update student successful"
// @Failure 400 {object} ErrorResponse "Invalid request body"
// @Failure 401 {object} ErrorResponse "Invalid credentials"
// @Failure 404 {object} ErrorResponse "Student not found"
// @Router /students/{id} [put]
func updateStudent(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(400).JSON(ErrorResponse{
			Success: false,
			Message: "Missing parameter!",
		})
	}

	var body UpdateStudentRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(ErrorResponse{
			Success: false,
			Message: "Body parsing failed!",
		})
	}

	var student models.Student

	// Update student
	for i := 0; i < len(students); i++ {
		s := &students[i] // Biar data master bisa berubah

		if id, _ := strconv.Atoi(id); students[i].ID == id {
			s.NIM = body.NIM
			s.Name = body.Name
			s.Email = body.Email
			s.Major = body.Major
			s.Semester = body.Semester

			student = *s

			break
		}
	}

	if student.ID == 0 {
		return c.Status(404).JSON(ErrorResponse{
			Success: false,
			Message: "Student not found!",
		})
	}

	return c.JSON(UpdateStudentResponse{
		Success: true,
		Message: "Update student success!",
		Data:    student,
	})
}

// DeleteStudent godoc
// @Summary Delete student
// @Description Delete student
// @Tags Students
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Student id"
// @Success 200 {object} DeleteStudentResponse "Update student successful"
// @Failure 401 {object} ErrorResponse "Invalid credentials"
// @Failure 404 {object} ErrorResponse "Student not found"
// @Router /students/{id} [delete]
func DeleteStudent(c *fiber.Ctx) error {
	id := c.Params("id")

	var student models.Student

	// Update student
	for i := 0; i < len(students); i++ {
		if id, _ := strconv.Atoi(id); students[i].ID == id {
			student = students[i]
			students = append(students[:i], students[i+1:]...)
			break
		}
	}

	if student.ID == 0 {
		return c.Status(404).JSON(ErrorResponse{
			Success: false,
			Message: "Student not found!",
		})
	}

	return c.JSON(DeleteStudentResponse{
		Success: true,
		Message: "Delete student success!",
		Data:    student,
	})
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
