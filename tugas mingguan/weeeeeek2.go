package main

import (
	"fmt"
)

func main() {
	var p, l int

	// Langkah 2: Input
	fmt.Print("Masukkan panjang (p): ")
	fmt.Scan(&p)
	fmt.Print("Masukkan lebar (l): ")
	fmt.Scan(&l)

	// Langkah 3: Proses
	L := p * l     // Menghitung luas
	K := 2*p + 2*l // Menghitung keliling

	// Langkah 4: Output
	fmt.Printf("Keliling (K): %d\n", K)
	fmt.Printf("Luas (L): %d\n", L)
}
