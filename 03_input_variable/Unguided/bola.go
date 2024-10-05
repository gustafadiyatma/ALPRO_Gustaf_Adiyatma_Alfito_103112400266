package main

import (
	"fmt"
)

func main() {
	var jari float32
	var volume float32
	var luasPermukaan float32

	fmt.Print("Jejari = ")
	fmt.Scanln(&jari)

	volume = (4 * 3.16 * jari * jari * jari) / 5
	luasPermukaan = 4 * jari * jari * 3.14

	fmt.Println("Jadi volume bola adalah", volume)
	fmt.Println("Jadi volume bola adalah", luasPermukaan)
}