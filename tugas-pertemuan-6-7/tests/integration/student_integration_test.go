package integration

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
	"tugas-pertemuan-7/models"
	"tugas-pertemuan-7/tests"
	"tugas-pertemuan-7/utils"

	"github.com/stretchr/testify/assert"
)

var user = models.User{
	ID:       1,
	Username: "admin",
	Password: "admin123",
	Role:     "admin",
}

func TestGetAllStudents(t *testing.T) {
	app := tests.SetupTestApp()

	req := httptest.NewRequest(
		"GET",
		"/api/students",
		nil,
	)

	token, _ := utils.CreateJwt(&user)

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var response models.GetStudentsResponse
	json.NewDecoder(resp.Body).Decode(&response)

	assert.True(t, response.Success)

	for _, student := range response.Data {
		assert.IsType(t, models.Student{}, student)
	}
}

func TestGetStudentByID(t *testing.T) {
	app := tests.SetupTestApp()

	req := httptest.NewRequest(
		"GET",
		"/api/students/1",
		nil,
	)

	token, _ := utils.CreateJwt(&user)

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var response models.GetStudentResponse
	json.NewDecoder(resp.Body).Decode(&response)

	assert.True(t, response.Success)
	assert.IsType(t, models.Student{}, response.Data)
}

func TestGetStudentByID_NotFound(t *testing.T) {
	app := tests.SetupTestApp()

	req := httptest.NewRequest(
		"GET",
		"/api/students/0",
		nil,
	)

	token, _ := utils.CreateJwt(&user)

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, 404, resp.StatusCode)

	var response models.GetStudentResponse
	json.NewDecoder(resp.Body).Decode(&response)

	assert.False(t, response.Success)
}

func TestCreateStudent(t *testing.T) {
	app := tests.SetupTestApp()

	studentData := `{"nim":"2021001","name":"Luthfi","email":"luthfi@example.com","major":"Computer Science","semester":1}`

	req := httptest.NewRequest(
		"POST",
		"/api/students",
		strings.NewReader(studentData),
	)

	token, _ := utils.CreateJwt(&user)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode)

	var response models.CreateStudentResponse
	json.NewDecoder(resp.Body).Decode(&response)

	assert.True(t, response.Success)
	assert.NotEmpty(t, response.Data)
	assert.IsType(t, models.Student{}, response.Data)
}

func TestUpdateStudent(t *testing.T) {
	app := tests.SetupTestApp()

	studentData := `{"nim":"2021001","name":"Luthfi Edited","email":"luthfi@example.com","major":"Computer Science","semester":1}`

	req := httptest.NewRequest(
		"PUT",
		"/api/students/1",
		strings.NewReader(studentData),
	)

	token, _ := utils.CreateJwt(&user)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode)

	var response models.UpdateStudentResponse
	json.NewDecoder(resp.Body).Decode(&response)

	assert.True(t, response.Success)
	assert.NotEmpty(t, response.Data)
	assert.IsType(t, models.Student{}, response.Data)
}
