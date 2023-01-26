package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	tNumbers := []int{}
	for _, val := range *numbers {
		tNumbers = append(tNumbers, transform(val))
	}
	return tNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
