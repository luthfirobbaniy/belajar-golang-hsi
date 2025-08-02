package worker

import (
	"fmt"
	"sync"
	"time"
	"tugas-pertemuan-4/models"

	"gorm.io/gorm"
)

type Result struct {
	MahasiswaID   uint
	MahasiswaName string
	TugasID       uint
	TugasJudul    string
	HasilNilai    int
}

func Assignment(
	db *gorm.DB,
	wg *sync.WaitGroup,
	mahasiswa models.Mahasiswa,
	jobs <-chan models.Tugas,
	result chan<- Result,
) {
	defer wg.Done()

	for job := range jobs {
		job.MahasiswaID = &mahasiswa.ID

		db.Save(job)

		result <- Result{
			MahasiswaID:   mahasiswa.ID,
			MahasiswaName: mahasiswa.Name,
			TugasJudul:    job.Judul,
			TugasID:       job.ID,
		}

		fmt.Printf(
			"Tugas '%s' diberikan ke %s\n",
			job.Judul,
			mahasiswa.Name,
		)

		// Delay to simulate processing delay
		time.Sleep(2 * time.Second)
	}
}
