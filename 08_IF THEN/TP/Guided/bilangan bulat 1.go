package main

import (
	"fmt"
	"math"
)

func main() {
	var number int

	// Meminta input dari pengguna
	fmt.Print("Masukkan sebuah bilangan bulat: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Input tidak valid, silakan masukkan bilangan bulat.")
		return
	}

	// Menghitung nilai absolut
	absoluteValue := int(math.Abs(float64(number)))

	// Menampilkan hasil
	fmt.Printf("Nilai absolut dari %d adalah %d\n", number, absoluteValue)
}
