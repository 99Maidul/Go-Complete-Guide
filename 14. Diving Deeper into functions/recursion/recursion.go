package recursion

import "fmt"

func main() {
	fact := factorial(5)
	fmt.Println(fact)
}

func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
	// result := 1
	// for i := 2; i <= number; i++ {
	// 	result *= i
	// }
	// return result
}
