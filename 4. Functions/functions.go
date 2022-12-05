package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a, b := generateRandomNumber()
	sum := add(a, b)
	pntNum(sum)
}
func add(x, y int) (sum int) {
	sum = x + y
	return
}

func pntNum(num int) {
	fmt.Println(num)
}

func generateRandomNumber() (r1 int, r2 int) {
	r1 = rand.Intn(10)
	r2 = rand.Intn(10)
	return
}
