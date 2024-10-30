package main

import (
	"fmt"
)

func main() {
	var number int

	// Meminta input dari pengguna
	fmt.Print("Masukkan suatu bilangan bulat: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Input tidak valid. Silakan masukkan bilangan bulat.")
		return
	}

	// Menentukan apakah bilangan positif atau bukan
	if number > 0 {
		fmt.Println("Bilangan bulat adalah positif.")
	} else {
		fmt.Println("Bilangan bulat adalah bukan positif.")
	}
