package main

import (
	"fmt"

	"tugas-pertemuan-3/mahasiswa"
)

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

	// Sum all students / mahasiswa age with closure
	ages := &mahasiswa.Ages{}
	sumAges := ages.Sum()
	sumAges(ali.GetUmur(), luthfi.GetUmur())

	// Print the total age
	fmt.Printf(
		"Total Umur Mahasiswa: %d",
		sumAges(0).GetTotal(), // Chaining function
	)
}
