package utils

import (
	"errors"
	"os"
	"tugas-pertemuan-7/models"

	"github.com/golang-jwt/jwt/v5"

	"time"

	_ "tugas-pertemuan-7/docs"
)

type Claims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func CreateJwt(data *models.User) (string, error) {
	claims := Claims{
		Id:       data.ID,
		Username: data.Username,
		Role:     data.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ParseJwt(authHeader string) (*Claims, error) {
	var tokenString string

	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		tokenString = authHeader[7:]
	}

	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(t *jwt.Token) (any, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		},
	)

	if err != nil || !token.Valid {
		return nil, errors.New("Invalid or expired token")
	}

	claims, ok := token.Claims.(*Claims)

	if !ok {
		return nil, errors.New("Invalid token claims")
	}

	return claims, nil
}
