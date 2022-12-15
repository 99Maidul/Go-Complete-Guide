package main

import "fmt"

func main() {
	age := 23
	fmt.Println(age)

	myAge := &age
	fmt.Println(*myAge)
	*myAge = 24
	fmt.Println(age)

	doubledAge := doubleAge(myAge)
	fmt.Println(doubledAge)
	fmt.Println(age)
}

func doubleAge(a *int) int {
	result := *a * 2
	return result
}
