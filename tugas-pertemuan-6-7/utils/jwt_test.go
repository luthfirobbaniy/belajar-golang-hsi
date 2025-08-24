package utils

import (
	"fmt"
	"testing"
	"tugas-pertemuan-7/models"
)

var user = models.User{
	ID:       1,
	Username: "admin",
	Role:     "admin",
}

func TestGenerateToken(t *testing.T) {
	token, err := CreateJwt(&user)

	if err != nil {
		t.Fatal("CreateJWT() return unexpected error!")
	}

	if token == "" {
		t.Error("CreateJWT() should return token, not empty string!")
	}
}

func TestValidateToken(t *testing.T) {
	token, err := CreateJwt(&user)

	if err != nil {
		t.Fatalf("CreateJWT() return unexpected error: %v", err)
	}

	claims, err := ParseJwt(fmt.Sprintf("Bearer %s", token))

	if err != nil {
		t.Fatalf("ParseJwt() return unexpected error: %v", err)
	}

	if claims.Id != user.ID {
		t.Errorf("claims.Id: %d; Expected: %d", claims.Id, user.ID)
	}
}

func TestInvalidToken(t *testing.T) {
	cases := []struct {
		name        string
		token       string
		expectedErr string
	}{
		{
			"expired token",
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGUiOiJhZG1pbiIsImV4cCI6MTc1NTk1MzE2OCwiaWF0IjoxNzU1OTUzMTY3fQ.33pYu7qti7lwAY0gS0RPVau-gFLTAOFoFFUg0q7IuJo",
			"Invalid or expired token",
		},
		{
			"not jwt",
			"token",
			"Invalid or expired token",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := ParseJwt(fmt.Sprintf("Bearer %s", c.token))

			if err == nil {
				t.Fatalf("ParseJwt() should return an error: %v", err)
			}

			if err.Error() != c.expectedErr {
				t.Errorf("ParseJwt() error: %s; Expected %s", err.Error(), c.expectedErr)
			}
		})
	}
}
