package models

import (
	"gorm.io/gorm"
)

type Mahasiswa struct {
	gorm.Model
	Name  string
	Tugas []Tugas
}

// Disable the auto-plural ("Mahasiswas" -> "Mahasiswa")
func (Mahasiswa) TableName() string {
	return "Mahasiswa"
}
