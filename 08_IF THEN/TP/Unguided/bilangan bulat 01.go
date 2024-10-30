package main

import (
	"fmt"
)

func main() {
	var number int

	// Meminta input dari pengguna
	fmt.Print("Masukkan bilangan bulat: ")
	_, err := fmt.Scanf("%d", &number)
	if err != nil {
		fmt.Println("Input tidak valid. Harap masukkan bilangan bulat.")
		return
	}

	// Mengecek apakah bilangan tersebut genap dan negatif
	if number < 0 && number%2 == 0 {
		fmt.Println("Bilangan adalah genap negatif.")
	} else {
		fmt.Println("Bilangan bukan genap negatif.")
	}
}
