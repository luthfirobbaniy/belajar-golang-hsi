package main

import (
	"tugas-pertemuan-7/handlers"
	"tugas-pertemuan-7/middleware"
	"tugas-pertemuan-7/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	_ "tugas-pertemuan-7/docs"
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
func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Default: 500
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			return c.Status(code).JSON(models.ErrorResponse{
				Success: false,
				Message: err.Error(),
			})
		},
	})

	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(recover.New())

	godotenv.Load()

	// Swager
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	// app.Get("/swagger/*", swagger.HandlerDefault) // Alternative: Swagger from Fiber

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
	app.Get("/api/profile", middleware.Jwt, handlers.GetProfile)

	app.Listen(":3000")
}
