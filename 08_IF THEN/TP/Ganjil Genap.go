package main

import (
	"fmt"
)

func main() {
	var angka int

	// Meminta pengguna untuk memasukkan angka
	fmt.Print("Masukkan sebuah angka: ")
	fmt.Scanln(&angka)

	// Memeriksa apakah angka tersebut genap atau ganjil
	if angka%2 == 0 {
		fmt.Println("Angka adalah Genap.")
	} else {
		fmt.Println("Angka adalah Ganjil.")
	}
}
