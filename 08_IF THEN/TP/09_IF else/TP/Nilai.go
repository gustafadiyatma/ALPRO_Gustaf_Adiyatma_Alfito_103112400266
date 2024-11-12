package main

import (
	"fmt"
)

func main() {
	var nilai float64

	// Meminta input nilai dari pengguna
	fmt.Print("Masukkan nilai mahasiswa: ")
	fmt.Scan(&nilai)

	// Mengklasifikasikan nilai
	var indeks string
	if nilai > 90 {
		indeks = "A"
	} else if nilai >= 80 {
		indeks = "AB"
	} else if nilai >= 70 {
		indeks = "B"
	} else {
		indeks = "C"
	}

	// Menampilkan hasil
	fmt.Printf("Indeks untuk nilai %.2f adalah %s\n", nilai, indeks)
}
