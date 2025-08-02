package worker

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"tugas-pertemuan-4/models"

	"gorm.io/gorm"
)

func Grading(
	db *gorm.DB,
	wg *sync.WaitGroup,
	jobs <-chan Result,
	result chan<- Result,
) {
	defer wg.Done()

	for job := range jobs {
		job.HasilNilai = int(rand.IntN(100))

		hasil := models.Hasil{
			TugasID: job.TugasID,
			Nilai:   job.HasilNilai,
		}

		db.Save(&hasil)

		result <- job

		fmt.Printf(
			"Nilai %d diberikan ke %s untuk tugas '%s'\n",
			job.HasilNilai,
			job.MahasiswaName,
			job.TugasJudul,
		)
	}
}
