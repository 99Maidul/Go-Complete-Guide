package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type User struct {
	firstName  string
	lastName   string
	birthdate  string
	promptedAt time.Time
}

func main() {

	// Create a new bufio.Reader to read from stdin
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user for their first name
	var newUser User
	firstName, err := getUserData(reader, "Enter your first name: ")
	checkError(err)

	// Prompt the user for their last name
	lastName, err := getUserData(reader, "Enter your last name: ")
	checkError(err)

	// Prompt the user for their birthdate
	birthdate, err := getUserData(reader, "Enter your birthdate (MM/DD/YY): ")
	checkError(err)

	// Get the current time
	promptedAt := time.Now()

	// Output the results
	newUser = User{
		firstName:  firstName,
		lastName:   lastName,
		birthdate:  birthdate,
		promptedAt: promptedAt,
	}
	fmt.Println(newUser)
}
func getUserData(reader *bufio.Reader, promptText string) (string, error) {
	fmt.Print(promptText)
	userInput, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimRight(userInput, "\n"), nil
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
}
