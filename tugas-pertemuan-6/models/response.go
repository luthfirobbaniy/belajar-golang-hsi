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

// @Description Error response
type ErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Invalid request body"`
}
