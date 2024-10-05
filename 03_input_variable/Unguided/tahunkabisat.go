package main

import "fmt"

// Fungsi untuk memeriksa apakah tahun kabisat
func isLeapYear(year int) bool {
    if year%400 == 0 {
        return true
    }
    if year%100 == 0 {
        return false
    }
    if year%4 == 0 {
        return true
    }
    return false
}

func main() {
    var year int
    fmt.Print("Tahun: ")
    fmt.Scan(&year)
    
    if isLeapYear(year) {
        fmt.Println("Kabisat: true")
    } else {
        fmt.Println("Kabisat: false")
    }
}