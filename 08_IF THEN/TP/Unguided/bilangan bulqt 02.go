package main

import (
	"fmt"
)

func main() {
	var x, y int

	// Meminta input dari pengguna
	fmt.Print("Masukkan bilangan bulat positif x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan bulat positif y: ")
	fmt.Scan(&y)

	// Menentukan apakah x adalah faktor dari y
	isFactorXY := y%x == 0
	// Menentukan apakah y adalah faktor dari x
	isFactorYX := x%y == 0

	// Mencetak hasil
	fmt.Println(isFactorXY) // Baris pertama
	fmt.Println(isFactorYX) // Baris kedua
}
