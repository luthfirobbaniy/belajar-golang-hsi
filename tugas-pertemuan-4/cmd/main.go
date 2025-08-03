package main

import (
	"fmt"
	"sync"
	"tugas-pertemuan-4/config"
	"tugas-pertemuan-4/models"
	"tugas-pertemuan-4/worker"
)

func main() {
	db, _ := config.DB()

	var mahasiswa []models.Mahasiswa
	var tugas []models.Tugas

	// Auto Migration
	db.AutoMigrate(
		&models.Mahasiswa{},
		&models.Tugas{},
		&models.Hasil{},
	)

	// Mahasiswa Seed (Masukin ke model?)
	mahasiswa = []models.Mahasiswa{
		{Name: "Andi Pratama"},
		{Name: "Budi Santoso"},
		{Name: "Citra Lestari"},
		{Name: "Dian Kusuma"},
		{Name: "Eka Sari"},
	}

	// Tugas Seed
	tugas = []models.Tugas{
		{
			Judul:     "Tugas Pemrograman Goroutine",
			Deskripsi: "Tugas Pemrograman Goroutine",
		},
		{
			Judul:     "Tugas Implementasi WaitGroup",
			Deskripsi: "Tugas Implementasi WaitGroup",
		},
		{
			Judul:     "Tugas Implementasi Mutex",
			Deskripsi: "Tugas Implementasi Mutex",
		},
		{
			Judul:     "Tugas Implementasi Channel",
			Deskripsi: "Tugas Implementasi Channel",
		},
		{
			Judul:     "Tugas Implementasi API",
			Deskripsi: "Tugas Implementasi API",
		},
	}

	// Create Seed Data
	db.Create(&mahasiswa)
	db.Create(&tugas)

	var wg sync.WaitGroup

	db.Find(&mahasiswa)
	db.Find(&tugas)

	tugasChan := make(chan models.Tugas, len(tugas))
	assignmentResult := make(chan worker.Result, len(tugas))

	for _, mahasiswa := range mahasiswa {
		wg.Add(1)

		go worker.Assignment(
			db,
			&wg,
			mahasiswa,
			tugasChan,
			assignmentResult,
		)
	}

	for _, tugas := range tugas {
		tugasChan <- tugas
	}

	close(tugasChan)

	wg.Wait()

	gradingResult := make(chan worker.Result, len(tugas))

	for range tugas {
		wg.Add(1)

		go worker.Grading(
			db,
			&wg,
			assignmentResult,
			gradingResult,
		)
	}

	close(assignmentResult)

	wg.Wait()

	fmt.Println("\nHasil Tugas Mahasiswa:")

	for range tugas {
		result := <-gradingResult

		fmt.Printf(
			"%s - %s: %d\n",
			result.MahasiswaName,
			result.TugasJudul,
			result.HasilNilai,
		)
	}
}
