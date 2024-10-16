package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scanln(&b)

	fmt.Println("Baris bilangan dari", a, "sampai dengan", b, "adalah:")
	for i := a; i <= b; i++ {
		fmt.Print(i, " ")
	}
}
