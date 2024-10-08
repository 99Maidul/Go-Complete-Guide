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

type PromptStorer interface {
	Storer
	Prompter
}
type userInputData struct {
	input string
}

type user struct {
	firstName string
	lastName  string
	*userInputData
}

func newUser(first, last string) *user {
	return &user{firstName: first, lastName: last, userInputData: &userInputData{}}
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

func (usr *userInputData) Store(fileName string) {
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
	// data.PromptForInput()
	// data.Store("user1.txt")
	name := newUser("John", "Doe")
	name.PromptForInput()
	name.Store("username.txt")

	handleUserInput(data)
}

func handleUserInput(container PromptStorer) {
	fmt.Println("Ready to store your data!")
	container.PromptForInput()
	container.Store("user2.txt")
}
