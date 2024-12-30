package main

import "fmt"

func main() {

	var (
		jumlah             int
		username, password string
		selesai            bool
	)

	jumlah = -1

	for selesai = false; !selesai; {
		fmt.Scan(&username, &password)

		jumlah++
		selesai = username == "admin" && password == "admin"
	}

	fmt.Println(jumlah)
	fmt.Print("Login berhasil")

}
