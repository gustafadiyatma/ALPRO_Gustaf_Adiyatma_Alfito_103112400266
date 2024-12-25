package main

import (
	"fmt"
)

func main() {
	var usia int
	var memilikiKTP bool

	// Membaca input usia
	fmt.Print("Masukkan usia: ")
	fmt.Scan(&usia)

	// Membaca input status kepemilikan KTP
	fmt.Print("Apakah memiliki KTP? (true/false): ")
	fmt.Scan(&memilikiKTP)

	// Menentukan apakah penduduk bisa membuat KTP
	if usia >= 17 && !memilikiKTP {
		fmt.Println("Anda memenuhi syarat untuk membuat KTP.")
	} else if usia < 17 {
		fmt.Println("Anda tidak memenuhi syarat karena usia kurang dari 17 tahun.")
	} else {
		fmt.Println("Anda sudah memiliki KTP, tidak perlu membuat KTP baru.")
	}
}
