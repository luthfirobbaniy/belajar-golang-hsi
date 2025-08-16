package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	_ "tugas-pertemuan-6/docs"
	"tugas-pertemuan-6/models"
)

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

var latestStudentId = 2

// getStudents godoc
// @Summary Get all student
// @Description Get all student data
// @Tags Students
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.GetStudentsResponse "Get students successful"
// @Failure 401 {object} models.ErrorResponse "Authorization header required or Invalid token"
// @Router /students [get]
func GetStudents(c *fiber.Ctx) error {
	return c.JSON(models.GetStudentsResponse{
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
// @Success 200 {object} models.GetStudentResponse "Get student successful"
// @Failure 400 {object} models.ErrorResponse "Missing parameter"
// @Failure 401 {object} models.ErrorResponse "Authorization header required or Invalid token"
// @Failure 404 {object} models.ErrorResponse "Student not found"
// @Router /students/{id} [get]
func GetStudent(c *fiber.Ctx) error {
	if c.Params("id") == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Missing parameter!")
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
		return fiber.NewError(fiber.StatusNotFound, "Student not found!")
	}

	return c.JSON(models.GetStudentResponse{
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
// @Param request body models.CreateStudentRequest true "Create student data"
// @Success 200 {object} models.CreateStudentResponse "Create student successful"
// @Failure 400 {object} models.ErrorResponse "Invalid request body"
// @Failure 401 {object} models.ErrorResponse "Invalid credentials"
// @Router /students [post]
func CreateStudent(c *fiber.Ctx) error {
	var body models.CreateStudentRequest

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Body parsing failed!")
	}

	latestStudentId += 1

	newStudent := models.Student{
		ID:       latestStudentId,
		NIM:      body.NIM,
		Name:     body.Name,
		Email:    body.Email,
		Major:    body.Major,
		Semester: body.Semester,
	}

	students = append(students, newStudent)

	return c.Status(201).JSON(models.CreateStudentResponse{
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
// @Param id path int true "Student id"
// @Param request body models.UpdateStudentRequest true "Update student data"
// @Success 200 {object} models.UpdateStudentResponse "Update student successful"
// @Failure 400 {object} models.ErrorResponse "Invalid request body"
// @Failure 401 {object} models.ErrorResponse "Invalid credentials"
// @Failure 404 {object} models.ErrorResponse "Student not found"
// @Router /students/{id} [put]
func UpdateStudent(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Missing parameter!")
	}

	var body models.UpdateStudentRequest

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Body parsing failed!")
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
		return fiber.NewError(fiber.StatusNotFound, "Student not found!")
	}

	return c.JSON(models.UpdateStudentResponse{
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
// @Success 200 {object} models.DeleteStudentResponse "Update student successful"
// @Failure 401 {object} models.ErrorResponse "Invalid credentials"
// @Failure 404 {object} models.ErrorResponse "Student not found"
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
		return fiber.NewError(fiber.StatusNotFound, "Student not found!")
	}

	return c.JSON(models.DeleteStudentResponse{
		Success: true,
		Message: "Delete student success!",
		Data:    student,
	})
}
