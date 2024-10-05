package main

import (
	"fmt"
)

func main() {

	var f float32
	var rumus float32

	fmt.Println("Masukan uang rupiah")
	fmt.Scanln(&f)
	rumus = f / 15000
	fmt.Println("Jadi jumlah dolar adalah", rumus)
}