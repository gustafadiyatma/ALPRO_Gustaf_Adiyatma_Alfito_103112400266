package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	// Meminta input dari pengguna
	fmt.Print("Masukkan sebuah huruf: ")
	fmt.Scan(&input)

	// Mengubah input menjadi huruf kapital untuk memudahkan pengecekan
	input = strings.ToUpper(input)

	// Mengecek apakah input adalah huruf vokal atau konsonan
	if len(input) == 1 { // Pastikan hanya satu karakter yang dimasukkan
		switch input {
		case "A", "I", "U", "E", "O":
			fmt.Println("Huruf Vokal")
		default:
			fmt.Println("Huruf Konsonan")
		}
	} else {
		fmt.Println("Input tidak valid, silakan masukkan satu huruf.")
	}
}
