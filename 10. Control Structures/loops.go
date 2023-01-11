package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	selectedChoice, err := getUserChoice()

	if err != nil {
		fmt.Println("Invalid choice")
		return
	}

	if selectedChoice == "1" {
		calculateSumUptoNumber()
	} else if selectedChoice == "2" {
		calculateFactorial()
	} else if selectedChoice == "3" {
		calculateSumOfManually()
	} else {
		calculateSumOfList()
	}
}

func calculateSumUptoNumber() {}
func calculateFactorial()     {}
func calculateSumOfManually() {}
func calculateSumOfList()     {}

func getUserChoice() (string, error) {
	fmt.Println("Please make your choice: ")
	fmt.Println("1) Add up all the numbers of to number X")
	fmt.Println("2) Build the factorial up to number X")
	fmt.Println("3) Sum up manualy entered numbers")
	fmt.Println("4) Sum up the list of entered numbers")

	userInput, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}
	userInput = strings.Replace(userInput, "\n", "", -1)
	userInput = strings.Replace(userInput, "\r", "", -1)

	if userInput == "1" ||
		userInput == "2" ||
		userInput == "3" ||
		userInput == "4" {
		return userInput, nil
	} else {
		return "", errors.New("INVAID INPUT")
	}
}
