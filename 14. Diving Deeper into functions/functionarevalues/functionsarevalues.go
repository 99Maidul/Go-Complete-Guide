package functionsarevalues

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(doubled)
	fmt.Println(tripled)

	transformfn1 := getTransformerFunction(&numbers)
	transformfn2 := getTransformerFunction(&moreNumbers)

	transformedNumbers := transformNumbers(&numbers, transformfn1)
	moretransformedNumbers := transformNumbers(&moreNumbers, transformfn2)

	fmt.Println(transformedNumbers)
	fmt.Println(moretransformedNumbers)
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	tNumbers := []int{}
	for _, val := range *numbers {
		tNumbers = append(tNumbers, transform(val))
	}
	return tNumbers
}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
