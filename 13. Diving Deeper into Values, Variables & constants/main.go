package main

import "fmt"

const userName = "Maidul"
const age = 66 / 3

const (
	inputAttack   = iota
	specialAttack = iota + 10
	inputHeal
)

func main() {
	fmt.Println(inputAttack, specialAttack, inputHeal)
	fmt.Println(userName, age)
}
