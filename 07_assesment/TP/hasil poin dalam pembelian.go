package main

import (
	"fmt"
)

func main() {
	var jumlahBarang int

	// Mengambil input dari pengguna
	fmt.Print("Masukkan jumlah barang yang dibeli: ")
	_, err := fmt.Scan(&jumlahBarang)
	if err != nil || jumlahBarang <= 0 {
		fmt.Println("Input tidak valid. Harap masukkan bilangan bulat positif.")
		return
	}

	// Menghitung total poin
	totalPoin := 0

	// Setiap barang memberi 10 poin
	totalPoin += jumlahBarang * 10

	// Jika lebih dari 5 barang, tambahkan 5 poin untuk setiap barang setelah barang kelima
	if jumlahBarang > 5 {
		totalPoin += (jumlahBarang - 5) * 5
	}

	// Menampilkan hasil
	fmt.Printf("Total poin yang didapatkan pelanggan: %d\n", totalPoin)
}
