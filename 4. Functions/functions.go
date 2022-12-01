package main

import (
	"fmt"
)

func main() {
	sum := add(1, 2)
	fmt.Println(sum)
}
func add(x, y int) int {
	return x + y
}
