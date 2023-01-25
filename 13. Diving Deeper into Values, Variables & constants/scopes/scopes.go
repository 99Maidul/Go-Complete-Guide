package scopes

import "fmt"

var userName = "Maidul"

func main() {
	shouldContinue := true
	if shouldContinue {
		userAge := 23

		fmt.Printf("My name is %v and I am %v years old.", userName, userAge)
	}
}
