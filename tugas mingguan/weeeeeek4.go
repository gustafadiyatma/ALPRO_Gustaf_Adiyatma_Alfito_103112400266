package main

import (
	"fmt"
)

// Definisikan tipe bentukan untuk menyimpan data transaksi
type Transaksi struct {
	NamaBarang  string
	Jumlah      int
	HargaSatuan float64
}

// Fungsi untuk menghitung total harga
func (t Transaksi) TotalHarga() float64 {
	return float64(t.Jumlah) * t.HargaSatuan
}

func main() {
	var namaBarang string
	var jumlah int
	var hargaSatuan float64

	// Minta input dari pengguna
	fmt.Print("Masukkan nama barang: ")
	fmt.Scanln(&namaBarang)
	fmt.Print("Masukkan jumlah: ")
	fmt.Scan(&jumlah)
	fmt.Print("Masukkan harga satuan: ")
	fmt.Scan(&hargaSatuan)

	// Buat objek transaksi
	transaksi := Transaksi{
		NamaBarang:  namaBarang,
		Jumlah:      jumlah,
		HargaSatuan: hargaSatuan,
	}

	// Hitung dan tampilkan total harga
	total := transaksi.TotalHarga()
	fmt.Printf("Total harga untuk %s adalah: %.2f\n", transaksi.NamaBarang, total)
}
