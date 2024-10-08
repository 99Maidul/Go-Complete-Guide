package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
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

func calculateSumUptoNumber() {
	fmt.Print("Please enter your number:")
	chosenNumber, err := getInputNumber()
	if err != nil {
		fmt.Println("Invalid Number Input")
		return
	}
	fmt.Println(chosenNumber)
	sum := 0
	for i := 1; i <= chosenNumber; i++ {
		sum = sum + i
	}
	fmt.Printf("Result is: %v", sum)
}
func calculateFactorial() {
	fmt.Print("Please enter your number:")
	chosenNumber, err := getInputNumber()
	if err != nil {
		fmt.Println("Invalid Number Input")
		return
	}
	factorial := 1
	for i := 1; i <= chosenNumber; i++ {
		factorial = factorial * i
	}
	fmt.Printf("Result is: %v", factorial)
}

func calculateSumOfManually() {
	isEnteringNumbers := true
	sum := 0
	fmt.Println("Keep on entering numbers, the program will stop when you enter a non-number")
	for isEnteringNumbers {
		chosenNumber, err := getInputNumber()
		sum = sum + chosenNumber

		isEnteringNumbers = err == nil
	}
	fmt.Printf("Result is: %v", sum)

}
func calculateSumOfList() {
	fmt.Println("Please enter a list of numbers separated by comma:xxx")
	inputNumberList, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Invalid Number Input")
		return
	}
	inputNumberList = strings.TrimSuffix(inputNumberList, "\n")
	inputNumberList = strings.TrimSuffix(inputNumberList, "\r")
	inputNumberList = strings.Replace(inputNumberList, "\n", "", -1)

	inputNumbers := strings.Split(inputNumberList, ",")

	sum := 0
	for index, value := range inputNumbers {
		fmt.Printf("Index: %v, Value: %v\n", index, value)
		number, err := strconv.ParseInt(value, 0, 64)
		if err != nil {
			continue
		}
		sum = sum + int(number)
	}
	fmt.Printf("Result is: %v\n", sum)
}

func getInputNumber() (int, error) {
	inputNumber, err := reader.ReadString('\n')
	inputNumber = strings.TrimSuffix(inputNumber, "\n")
	inputNumber = strings.TrimSuffix(inputNumber, "\r")
	if err != nil {

		return 0, err
	}
	inputNumber = strings.Replace(inputNumber, "\n", "", -1)
	chosenNumber, err := strconv.ParseInt(inputNumber, 0, 64)
	if err != nil {

		return 0, err
	}
	return int(chosenNumber), nil
}
func getUserChoice() (string, error) {
	fmt.Println("Please make your choice: ")
	fmt.Println("1) Add up all the numbers of to number X")
	fmt.Println("2) Build the factorial up to number X")
	fmt.Println("3) Sum up manualy entered numbers")
	fmt.Println("4) Sum up the list of entered numbers")

	fmt.Println("Please make your choice: ")
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
