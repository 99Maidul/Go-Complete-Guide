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
	PromptForinput()
}

type userInputData struct {
	input string
}

func newUserInputData() *userInputData {
	return &userInputData{}
}

func (usr *userInputData) PromptForinput() {
	fmt.Print("Your Input Please ")
	reader := bufio.NewReader(os.Stdin)

	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Fetching User Input Failed")
		return
	}
	usr.input = userInput
}
