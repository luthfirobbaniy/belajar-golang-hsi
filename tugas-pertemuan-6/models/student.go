package models

type Student struct {
	ID       int    `json:"id"`
	NIM      string `json:"nim"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Major    string `json:"major"`
	Semester int    `json:"semester"`
}
