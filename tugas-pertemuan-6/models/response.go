package models

// @Description Login request payload
type LoginRequest struct {
	Username string `json:"username" example:"luthfi"`
	Password string `json:"password" example:"123"`
}

// @Description Login response
type LoginResponse struct {
	Success bool      `json:"success" example:"true"`
	Message string    `json:"message" example:"Login successful!"`
	Data    LoginData `json:"data"`
}

// @Description Login response data
type LoginData struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Imx1dGhmaSIsImV4cCI6MTc1NTMzMTU5NiwiaWF0IjoxNzU1MjQ1MTk2fQ.7WktpMm0AyyfXUR5x68Om7Pps9uR1resDlh2bz9C_J8"`
}

// @Description Register response data
type RegisterRequest struct {
	Username string `json:"username" example:"luthfi"`
	Password string `json:"password" example:"123"`
}

// @Description Register response data
type RegisterResponse struct {
	Success bool      `json:"success" example:"true"`
	Message string    `json:"message" example:"Login successful!"`
	Data    LoginData `json:"data"`
}

// @Description Get students response
type GetStudentsResponse struct {
	Success bool      `json:"success" example:"true"`
	Message string    `json:"message" example:"Get students successful!"`
	Data    []Student `json:"data"`
}

// @Description Get student response
type GetStudentResponse struct {
	Success bool    `json:"success" example:"true"`
	Message string  `json:"message" example:"Get student successful!"`
	Data    Student `json:"data"`
}

// @Description Create student request
type CreateStudentRequest struct {
	NIM      string `json:"nim" example:"2021003"`
	Name     string `json:"name" example:"Budi"`
	Email    string `json:"email" example:"budi@example.com"`
	Major    string `json:"major" example:"Teknik Industri"`
	Semester int    `json:"semester" example:"2"`
}

// @Description Create student response
type CreateStudentResponse struct {
	Success bool    `json:"success" example:"true"`
	Message string  `json:"message" example:"Create student success!"`
	Data    Student `json:"data"`
}

// @Description Update student request
type UpdateStudentRequest struct {
	NIM      string `json:"nim" example:"2021001"`
	Name     string `json:"name" example:"Luthfi Edited"`
	Email    string `json:"email" example:"luthfi@example.com"`
	Major    string `json:"major" example:"Computer Science"`
	Semester int    `json:"semester" example:"1"`
}

// @Description Update student response
type UpdateStudentResponse struct {
	Success bool    `json:"success" example:"true"`
	Message string  `json:"message" example:"Update student success!"`
	Data    Student `json:"data"`
}

// @Description Delete student response
type DeleteStudentResponse struct {
	Success bool    `json:"success" example:"true"`
	Message string  `json:"message" example:"Delete student success!"`
	Data    Student `json:"data"`
}

// @Description Error response
type ErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Invalid request body"`
}
