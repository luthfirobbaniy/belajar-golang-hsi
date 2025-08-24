package tests

import (
	"tugas-pertemuan-7/handlers"
	"tugas-pertemuan-7/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupTestApp() *fiber.App {
	app := fiber.New()

	app.Post("/api/auth/login", handlers.Login)

	app.Get("/api/students", middleware.Jwt, handlers.GetStudents)
	app.Get("/api/students/:id", middleware.Jwt, handlers.GetStudent)
	app.Post("/api/students", middleware.Jwt, handlers.CreateStudent)
	app.Put("/api/students/:id", middleware.Jwt, handlers.CreateStudent)

	return app
}
