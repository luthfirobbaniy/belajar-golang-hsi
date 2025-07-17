# Soal 1 CLI Validasi (Tugas Pekan 2)

This project is my answer for **Soal 1 CLI Validasi (Tugas Pekan 2)** where we (the students) are asked to create a CLI input and validate the input.

## How To Play

Run file using `go run main.go` and the terminal will give you an output like this

```
Nama:
```

Input your `Nama` and press enter. You will see the terminal giving you another output like this

```
Nama: Luthfi
Umur: 
```

Input your `Umur` and press enter.

If your `Umur` is greater than or equal to 18, you will see an output like this:

```
Nama: Luthfi
Umur: 18
Error: Selamat datang Luthfi
```

Otherwise the output will be like this:

```
Nama: Luthfi
Umur: 17
Error: umur tidak valid (minimal 18 tahun)
```
