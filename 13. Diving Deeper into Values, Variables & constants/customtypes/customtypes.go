package customtypes

import "fmt"

type person struct {
	name string
	age  int
}

type PersonData map[string]person
type CustomNumber int

func (number CustomNumber) pow(power int) CustomNumber {
	result := number
	for i := 1; i < power; i++ {
		result = result * number
	}
	return CustomNumber(result)
}

func main() {
	var people PersonData = PersonData{
		"p1": {"Maidul", 23},
	}
	fmt.Println(people)

	var dummyNumber CustomNumber = 5
	poweredNumber := dummyNumber.pow(3)
	fmt.Println(poweredNumber)
}
