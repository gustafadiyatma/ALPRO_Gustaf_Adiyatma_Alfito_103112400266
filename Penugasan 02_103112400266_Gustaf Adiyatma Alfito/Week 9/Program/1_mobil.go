package main

import "fmt"

func main() {

	var orang int

	fmt.Print("Masukkan jumlah orang : ")
	fmt.Scan(&orang)

	mobil := orang / 7
	sisaKursi := orang % 7

	if sisaKursi == 0 {
		mobil = mobil
	} else {
		mobil = mobil + 1
	}

	fmt.Print(mobil)

}
