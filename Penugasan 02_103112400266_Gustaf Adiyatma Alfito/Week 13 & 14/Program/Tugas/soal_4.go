package main

import "fmt"

func main() {

	var angka, terbesar, digit int

	fmt.Scan(&angka)

	terbesar = 0

	for angka > 0 {

		digit = angka % 10

		if digit > terbesar {
			terbesar = digit
		}

		angka /= 10

	}

	fmt.Println(terbesar)

}
