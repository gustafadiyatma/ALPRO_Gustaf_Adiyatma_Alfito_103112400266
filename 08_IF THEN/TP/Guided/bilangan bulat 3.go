package main

import (
	"fmt"
)

func isEvenNegative(n int) bool {
	return n < 0 && n%2 == 0
}

func main() {
	var number int
	fmt.Print("Masukkan bilangan bulat: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Input tidak valid.")
		return
	}

	result := isEvenNegative(number)
	fmt.Println("Apakah bilangan tersebut genap negatif?", result)
}
