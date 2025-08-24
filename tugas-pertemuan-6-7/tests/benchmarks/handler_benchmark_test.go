package benchmark

import (
	"net/http/httptest"
	"testing"
	"tugas-pertemuan-7/models"
	"tugas-pertemuan-7/tests"
	"tugas-pertemuan-7/utils"
)

var user = models.User{
	ID:       1,
	Username: "admin",
	Password: "admin123",
	Role:     "admin",
}

func BenchmarkGetStudents(b *testing.B) {
	app := tests.SetupTestApp()

	for b.Loop() {
		req := httptest.NewRequest("GET", "/api/students", nil)

		token, _ := utils.CreateJwt(&user)
		req.Header.Set("Authorization", "Bearer "+token)

		resp, _ := app.Test(req)
		resp.Body.Close()
	}
}

func BenchmarkGenerateToken(b *testing.B) {
	for b.Loop() {
		utils.CreateJwt(&user)
	}
}
