package main

import (
	"fmt"
	"unicode"
)

func main() {
	var karakter rune
	fmt.Print("Masukkan satu karakter: ")
	fmt.Scanf("%c", &karakter)

	// Mengecek apakah karakter adalah huruf
	if unicode.IsLetter(karakter) {
		// Mengecek apakah karakter adalah vokal atau konsonan
		if karakter == 'a' || karakter == 'e' || karakter == 'i' || karakter == 'o' || karakter == 'u' ||
			karakter == 'A' || karakter == 'E' || karakter == 'I' || karakter == 'O' || karakter == 'U' {
			fmt.Println("Karakter tersebut adalah huruf vokal.")
		} else {
			fmt.Println("Karakter tersebut adalah huruf konsonan.")
		}
	} else {
		fmt.Println("Karakter tersebut bukan huruf.")
	}
}
