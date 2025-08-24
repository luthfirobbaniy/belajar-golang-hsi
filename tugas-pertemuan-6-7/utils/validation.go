package utils

import (
	"net/mail"
)

func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func ValidateNIM(nim string) bool {
	return nim != ""
}

func ValidateSemester(semester int) bool {
	return semester > 0
}
