package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	transformedNumbers := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	fmt.Println(transformedNumbers)

}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	tNumbers := []int{}
	for _, val := range *numbers {
		tNumbers = append(tNumbers, transform(val))
	}
	return tNumbers
}
