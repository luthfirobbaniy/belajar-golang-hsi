package main

import (
	"fmt"

	"tugas-pertemuan-3/mahasiswa"
)

type Ages struct {
	Total int
}

func (a *Ages) count() func(ages ...int) *Ages {
	return func(ages ...int) *Ages {
		for _, age := range ages {
			a.Total += age
		}
		return a
	}
}

func (a *Ages) getTotal() int {
	return a.Total
}

func main() {
	// Create students / mahasiswa
	ali := mahasiswa.BuatMahasiswa("Ali", 18, 100, 70)
	luthfi := mahasiswa.BuatMahasiswa("Luthfi", 24, 80, 100)

	// List of students / mahasiswa
	students := []*mahasiswa.Mahasiswa{ali, luthfi}

	// Print student information
	for _, student := range students {
		fmt.Println(student.Info())
		fmt.Printf(
			"Rata-rata nilai: %.2f \n--- \n",
			student.RataRata(student.Nilai...),
		)
	}

	// Print hardcoded data
	fmt.Printf("Nilai Maksimum: %d \n", mahasiswa.GetMaxNilai())
	fmt.Printf("Versi Package: %s \n", mahasiswa.Versi)

	// Count all students / mahasiswa age with closure
	ages := &Ages{}
	ageCounter := ages.count()
	ageCounter(ali.GetUmur(), luthfi.GetUmur())

	// Print the total age
	fmt.Printf(
		"Total Umur Mahasiswa: %d",
		ageCounter(0).getTotal(), // Chaining function
	)
}
