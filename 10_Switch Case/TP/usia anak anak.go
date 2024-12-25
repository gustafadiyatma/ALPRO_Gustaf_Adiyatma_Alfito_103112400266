package main

import (
	"fmt"
)

func main() {
	// Meminta pengguna untuk memasukkan usia
	var usia int
	fmt.Print("Masukkan usia Anda: ")
	_, err := fmt.Scan(&usia)

	if err != nil {
		fmt.Println("Input tidak valid, silakan masukkan angka untuk usia.")
		return
	}

	// Menggunakan switch untuk menentukan kategori usia
	switch {
	case usia >= 0 && usia <= 12:
		fmt.Println("Kategori Usia: Anak-anak (0 - 12 tahun)")
	case usia >= 13 && usia <= 17:
		fmt.Println("Kategori Usia: Remaja (13 - 17 tahun)")
	case usia >= 18 && usia <= 64:
		fmt.Println("Kategori Usia: Dewasa (18 - 64 tahun)")
	case usia >= 65:
		fmt.Println("Kategori Usia: Lansia (65 tahun ke atas)")
	default:
		fmt.Println("Usia tidak valid")
	}
}
