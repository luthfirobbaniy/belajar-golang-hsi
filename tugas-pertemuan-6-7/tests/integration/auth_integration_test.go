package integration

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"
	"tugas-pertemuan-7/models"
	"tugas-pertemuan-7/tests"

	"github.com/stretchr/testify/assert"
)

func TestLoginSuccess(t *testing.T) {
	app := tests.SetupTestApp()

	loginData := `{"username":"admin","password":"admin123"}`

	req := httptest.NewRequest(
		"POST",
		"/api/auth/login",
		strings.NewReader(loginData),
	)

	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var response models.LoginResponse
	json.NewDecoder(resp.Body).Decode(&response)

	// assert.True(t, response["success"].(bool))
	// assert.Contains(t, response, "token")
	assert.True(t, response.Success)
	assert.NotEmpty(t, response.Data.Token)
}

func TestLoginFailed(t *testing.T) {
	app := tests.SetupTestApp()

	loginData := `{"username":"user","password":"user123"}`

	req := httptest.NewRequest(
		"POST",
		"/api/auth/login",
		strings.NewReader(loginData),
	)

	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode)

	var response models.LoginResponse
	json.NewDecoder(resp.Body).Decode(&response)

	assert.False(t, response.Success)
	assert.Empty(t, response.Data.Token)
}
