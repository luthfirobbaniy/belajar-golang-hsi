package models

import (
	"gorm.io/gorm"
)

type Hasil struct {
	gorm.Model
	TugasID uint `gorm:"unique"`
	Nilai   int
}

// Disable the auto-plural ("Hasils" -> "Hasil")
func (Hasil) TableName() string {
	return "Hasil"
}
