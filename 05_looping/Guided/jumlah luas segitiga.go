package main

import "fmt"

func calculateTriangleArea(base float64, height float64) float64 {
	return 0.5 * base * height
}

func main() {
	var n int
	fmt.Print("Enter the number of triangles: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var base float64
		var height float64
		fmt.Printf("Enter base of triangle %d: ", i+1)
		fmt.Scan(&base)
		fmt.Printf("Enter height of triangle %d: ", i+1)
		fmt.Scan(&height)

		area := calculateTriangleArea(base, height)
		fmt.Printf("The area of triangle %d is: %.2f\n", i+1, area)
	}
}
