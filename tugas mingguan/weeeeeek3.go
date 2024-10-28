package main

import (
	"fmt"
)

func main() {
	var number int

	// Meminta pengguna untuk memasukkan bilangan bulat positif
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scanln(&number)

	// Mengecek apakah bilangan tersebut ganjil
	isOdd := number%2 != 0

	// Menampilkan hasil
	fmt.Println(isOdd)
}
