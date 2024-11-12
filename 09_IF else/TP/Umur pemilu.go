package main

import (
	"fmt"
)

func main() {
	var umur int
	var kewarganegaraan string

	// Meminta input umur dari pengguna
	fmt.Print("Masukkan umur Anda: ")
	fmt.Scan(&umur)

	// Meminta input kewarganegaraan dari pengguna
	fmt.Print("Masukkan kewarganegaraan Anda (WNI/WNA): ")
	fmt.Scan(&kewarganegaraan)

	// Menentukan apakah pengguna bisa mengikuti pemilu
	if umur >= 17 && kewarganegaraan == "WNI" {
		fmt.Println("Anda bisa mengikuti pemilu.")
	} else {
		if umur < 17 {
			fmt.Println("Anda tidak bisa mengikuti pemilu karena umur Anda kurang dari 17 tahun.")
		}
		if kewarganegaraan != "WNI" {
			fmt.Println("Anda tidak bisa mengikuti pemilu karena Anda bukan WNI.")
		}
	}
}
