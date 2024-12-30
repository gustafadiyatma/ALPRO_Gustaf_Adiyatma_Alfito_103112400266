package main

import "fmt"

func main() {

	var kapasitas, volume, hasil int
	var selesai bool

	fmt.Scan(&kapasitas)

	for selesai = false; !selesai; {

		fmt.Scan(&volume)
		hasil += volume

		selesai = hasil >= kapasitas
		fmt.Println(selesai)

	}

}
