package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	firstName := getUserData("Enter your first name: ")
	lastName := getUserData("Enter your last name:")
	birthdate := getUserData("Enter your birthdate (MM/DD/YY): ")
	created := time.Now()

	// ... do something awesome with the gathered data
	fmt.Println(firstName, lastName, birthdate, created)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	userInput, _ := reader.ReadString('\n')

	cleanedInput := strings.Replace(userInput, "\n", "", -1)

	return cleanedInput
}
