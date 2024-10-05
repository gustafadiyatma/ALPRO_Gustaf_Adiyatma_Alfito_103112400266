package main

import "fmt"

func main() {
    var celsius float64
    fmt.Print("Temperatur Celsius: ")
    fmt.Scan(&celsius)

    // Konversi dari Celsius ke Reamur, Fahrenheit, dan Kelvin
    reamur := celsius * 4 / 5
    fahrenheit := (celsius * 9 / 5) + 32
    kelvin := celsius + 273.15

    // Output hasil konversi
    fmt.Printf("Derajat Reamur: %.2f\n", reamur)
    fmt.Printf("Derajat Fahrenheit: %.2f\n", fahrenheit)
    fmt.Printf("Derajat Kelvin: %.2f\n", kelvin)
}