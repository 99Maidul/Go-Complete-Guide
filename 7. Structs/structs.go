package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type User struct {
	firstName  string
	lastName   string
	birthdate  string
	promptedAt string
}

func NewUser(fName string, lName string, bDate string) User {
	promptedAt := time.Now().Format("2006-01-02 15:04:05")
	user := User{
		firstName:  fName,
		lastName:   lName,
		birthdate:  bDate,
		promptedAt: promptedAt,
	}
	return user
}

func (u User) String() string {
	return fmt.Sprintf("First name: %s\nLast name: %s\nBirthdate: %s\nPrompted at: %s", u.firstName, u.lastName, u.birthdate, u.promptedAt)
}

func getUserData(reader *bufio.Reader, prompt string) (string, error) {
	// Print the prompt
	fmt.Print(prompt)

	// Read the input
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	// Remove the trailing newline character
	input = input[:len(input)-1]

	return input, nil
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	// Create a new bufio.Reader to read from stdin
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user for their first name
	firstName, err := getUserData(reader, "Enter your first name: ")
	checkError(err)

	// Prompt the user for their last name
	lastName, err := getUserData(reader, "Enter your last name: ")
	checkError(err)

	// Prompt the user for their birthdate
	birthdate, err := getUserData(reader, "Enter your birthdate (MM/DD/YY): ")
	checkError(err)

	// Create a new user with the collected data
	user := NewUser(firstName, lastName, birthdate)

	// Output the user data
	fmt.Println(user)
}
