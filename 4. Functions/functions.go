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
func add(x, y int) int {
	return x + y
}

func pntNum(num int) {
	fmt.Println(num)
}

func generateRandomNumber() (int, int) {
	randomNumber1 := rand.Intn(10)
	randomNumber2 := rand.Intn(10)
	return randomNumber1, randomNumber2
}
