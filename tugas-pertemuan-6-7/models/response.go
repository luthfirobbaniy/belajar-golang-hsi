package models

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

// @Description Create student response
type CreateStudentResponse struct {
	Success bool    `json:"success" example:"true"`
	Message string  `json:"message" example:"Create student success!"`
	Data    Student `json:"data"`
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

// @Description Get profile response
type GetProfileResponse struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message" example:"Get profile success!"`
	Data    ProfileData `json:"data"`
}

// @Description Get profile response data
type ProfileData struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

// @Description Error response
type ErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Invalid request body"`
}
