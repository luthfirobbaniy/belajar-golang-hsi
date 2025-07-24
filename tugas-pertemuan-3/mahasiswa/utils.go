package mahasiswa

var maxNilai int = 100
var Versi string = "v1.0.0"

func hitungRataRata(nilai ...int) float64 {
	var total int

	for _, item := range nilai {
		total += item
	}

	return float64(total / len(nilai))
}

func BuatMahasiswa(nama string, umur int, nilai ...int) *Mahasiswa {
	return &Mahasiswa{
		umur:     umur,
		Nama:     nama,
		Nilai:    nilai,
		nilaiAvg: hitungRataRata(nilai...),
	}
}

func GetMaxNilai() int {
	return maxNilai
}
