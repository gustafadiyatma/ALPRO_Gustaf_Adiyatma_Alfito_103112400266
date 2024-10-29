package main

import (
	"fmt"
)

func main() {
	var nilai int

	// Meminta pengguna untuk memasukkan nilai ujian
	fmt.Print("Masukkan nilai ujian siswa: ")
	_, err := fmt.Scan(&nilai)
	if err != nil {
		fmt.Println("Input tidak valid. Silakan masukkan angka.")
		return
	}

	// Memeriksa apakah siswa lulus atau tidak
	if nilai >= 70 {
		fmt.Println("Lulus")
	} else {
		fmt.Println("Tidak Lulus")
	}
}
