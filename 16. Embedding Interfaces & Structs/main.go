package main

import (
	"bufio"
	"fmt"
	"os"
)

type Storer interface {
	Store(fileName string)
}

type Prompter interface {
	PromptForInput()
}

type userInputData struct {
	input string
}

func newUserInputData() *userInputData {
	return &userInputData{}
}

func (usr *userInputData) PromptForInput() {
	fmt.Print("Your Input Please: ")
	reader := bufio.NewReader(os.Stdin)

	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Fetching User Input Failed")
		return
	}
	usr.input = userInput
}

func (usr *userInputData) store(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Creating File Failed")
		return
	}
	defer file.Close()
	file.WriteString(usr.input)
}

func main() {
	data := newUserInputData()
	data.PromptForInput()
	data.store("user1.txt")
	// handleUserInput(data)
}

// func handleUserInput(container Prompter) {
// 	fmt.Println("Resdy to store your data!")
// 	container.PromptForInput()
// 	container.Store("user1.txt")
// }
