package models

import (
	"gorm.io/gorm"
)

type Tugas struct {
	gorm.Model
	Judul       string
	Deskripsi   string
	MahasiswaID *uint
	Hasil       Hasil
}

// Disable the auto-plural ("Tugass" -> "Tugas")
func (Tugas) TableName() string {
	return "Tugas"
}
