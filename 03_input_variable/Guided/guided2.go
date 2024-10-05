package main

import (
	"fmt"
)

func main() {

	var f float32
	var rumus float32

	fmt.Println("Masukan angka untuk x")
	fmt.Scanln(&f)
	rumus = (2/(f+5) + 5)
	fmt.Println("Jadi hasilnya adalah", rumus)
}