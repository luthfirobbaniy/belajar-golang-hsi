package models

type User struct {
	ID       int    `json:"id" example:"1"`
	Username string `json:"username" example:"luthfi"`
	Password string `json:"password" example:"123"`
	Role     string `json:"role" example:"admin"`
}
