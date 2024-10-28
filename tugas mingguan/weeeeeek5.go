package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah bilangan: ")
	fmt.Scan(&n)

	// Membuat slice untuk menyimpan bilangan
	numbers := make([]int, n)

	fmt.Println("Masukkan", n, "bilangan:")
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	// Menghitung hasil perkalian dengan indeks
	results := make([]int, n)
	for i := 0; i < n; i++ {
		results[i] = numbers[i] * (i + 1) // Indeks dimulai dari 1
	}

	// Menampilkan hasil
	fmt.Println("Hasil perkalian setiap bilangan dengan indeksnya:")
	for _, result := range results {
		fmt.Println(result)
	}
}
