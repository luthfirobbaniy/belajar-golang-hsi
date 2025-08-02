package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DB() (*gorm.DB, error) {
	dsn := "postgres://postgres:postgres@localhost:5432/mahasiswa_db"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
