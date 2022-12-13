package main

import "fmt"

func main() {
	pi := 3.14
	radius := 5

	var circumference float64 = 2 * pi * float64(radius)
	fmt.Printf("For a circle with a radius of %v the circumference is %.2f", radius, circumference)
}
