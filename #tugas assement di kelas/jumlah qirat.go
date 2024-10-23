package main

import "fmt"

func main() {
	var qirat int
	fmt.Print("Masukkan jumlah qirat: ")
	fmt.Scanln(&qirat)

	// 1 fals = 6 qirat
	fals := qirat / 6
	sisaQirat := qirat % 6

	// 1 dirham = 10 fals
	dirham := fals / 10
	sisaFals := fals % 10

	// 1 dinar = 10 dirham
	dinar := dirham / 10
	sisaDirham := dirham % 10

	fmt.Printf("Dinar: %d\n", dinar)
	fmt.Printf("Dirham: %d\n", sisaDirham)
	fmt.Printf("Fals: %d\n", sisaFals)
	fmt.Printf("Qirat: %d\n", sisaQirat)
}
