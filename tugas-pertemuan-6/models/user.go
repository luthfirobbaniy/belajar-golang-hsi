package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"` // Di real app, harus di-hash
	Role     string `json:"role"`     // "admin" atau "student"
}
