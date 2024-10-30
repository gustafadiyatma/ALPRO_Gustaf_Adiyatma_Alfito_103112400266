package main

import (
	"fmt"
	"math"
)

func main() {
	var jumlahOrang int

	// Meminta input dari pengguna
	fmt.Print("Masukkan jumlah orang yang akan melakukan touring: ")
	fmt.Scan(&jumlahOrang)

	// Menghitung jumlah motor yang diperlukan
	jumlahMotor := int(math.Ceil(float64(jumlahOrang) / 2))

	// Menampilkan hasil
	fmt.Printf("Jumlah motor yang diperlukan: %d\n", jumlahMotor)
}
