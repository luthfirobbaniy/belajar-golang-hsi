package mahasiswa

import (
	"fmt"
)

type Mahasiswa struct {
	Nama     string
	Nilai    []int
	umur     int
	nilaiAvg float64
}

type Deskripsi interface {
	Info() string
	RataRata() float64
	GetUmur() int
}

func (m Mahasiswa) Info() string {
	return fmt.Sprintf(
		"Nama: %s, Umur: %d",
		m.Nama,
		m.GetUmur(),
	)
}

func (m Mahasiswa) RataRata(nilai ...int) float64 {
	return hitungRataRata(nilai...)
}

func (m Mahasiswa) GetUmur() int {
	return m.umur
}
