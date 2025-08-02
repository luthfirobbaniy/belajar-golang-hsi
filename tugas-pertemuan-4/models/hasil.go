package models

import (
	"gorm.io/gorm"
)

type Hasil struct {
	gorm.Model
	TugasID uint
	Nilai   int
}

// Disable the auto-plural ("Hasils" -> "Hasil")
func (Hasil) TableName() string {
	return "Hasil"
}
