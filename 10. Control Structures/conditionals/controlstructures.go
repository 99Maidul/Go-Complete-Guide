package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	userAge, err := getUserAge()
	if err != nil {
		fmt.Println("invalid input!")
		return
	}

	// isOldEnough := userAge >= 32
	if (userAge >= 30 && userAge < 50) || userAge >= 60 {
		fmt.Println("You are eligible for the senior job!")
	} else if userAge >= 50 {
		fmt.Println("The best age!")
	} else if userAge >= 18 {
		fmt.Println("You are a young adult adult")
	} else {
		fmt.Println("Sorry, you are not old enough")
	}
	fmt.Println("Goodbye")
}
func getUserAge() (int, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter your age: ")
	userAgeInput, _ := reader.ReadString('\n')
	userAgeInput = strings.TrimSuffix(userAgeInput, "\n")
	userAgeInput = strings.TrimSuffix(userAgeInput, "\r")
	userAgeInput = strings.Replace(userAgeInput, "\n", "", -1)
	userAge, err := strconv.ParseInt(userAgeInput, 0, 64)

	return int(userAge), err
}
