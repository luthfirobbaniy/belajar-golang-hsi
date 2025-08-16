package models

type Student struct {
	ID       int    `json:"id" example:"1"`
	NIM      string `json:"nim" example:"2021001"`
	Name     string `json:"name" example:"Luthfi"`
	Email    string `json:"email" example:"luthfi@example.com"`
	Major    string `json:"major" example:"Computer Science"`
	Semester int    `json:"semester" example:"1"`
}
