package models

type Student struct {
	ID       int    `json:"id"`
	NIM      string `json:"nim" example:"2021003"`
	Name     string `json:"name" example:"Budi"`
	Email    string `json:"email" example:"budi@example.com"`
	Major    string `json:"major" example:"Teknik Industri"`
	Semester int    `json:"semester" example:"2"`
}
