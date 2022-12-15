package main

import "fmt"

func main() {
	age := 23
	fmt.Println(age)

	myAge := &age
	fmt.Println(*myAge)
	*myAge = 24
	fmt.Println(age)
}
