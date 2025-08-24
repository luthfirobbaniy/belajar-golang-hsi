package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmail(t *testing.T) {
	cases := []struct {
		name     string
		email    string
		expected bool
	}{
		{"valid email", "student@hsi-sandbox.ac.id", true},
		{"invalid email", "invalid-email", false},
		{"empty email", "", false},
		{"email without domain", "student@", false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := ValidateEmail(c.email)
			assert.Equal(t, c.expected, result)
		})
	}
}

func TestValidateNIM(t *testing.T) {
	cases := []struct {
		name     string
		nim      string
		expected bool
	}{
		{"valid NIM", "2021001", true},
		{"empty NIM", "", false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := ValidateNIM(c.nim)
			assert.Equal(t, c.expected, result)
		})
	}
}

func TestValidateSemester(t *testing.T) {
	cases := []struct {
		name     string
		semester int
		expected bool
	}{
		{"semester is greater than 0", 1, true},
		{"semester is equal to 0", 0, false},
		{"semester is less than 0", -1, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := ValidateSemester(c.semester)
			assert.Equal(t, c.expected, result)
		})
	}
}
