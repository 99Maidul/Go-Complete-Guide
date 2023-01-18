package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetPlayerChoice(specialAtttackIsAvailable bool) string {

	for {
		playerChoice, _ := getPlayerInput()

		if playerChoice == "1" {
			return "ATTACK"
		} else if playerChoice == "2" {
			return "HEAL"
		} else if playerChoice == "3" && specialAtttackIsAvailable {
			return "SPECIAL_ATTACK"
		}
		fmt.Println("Fetching the user input failed. Please try again.")
	}
}

func getPlayerInput() (string, error) {
	fmt.Print("Your choice: ")

	userInput, err := reader.ReadString('\n')
	userInput = strings.TrimSuffix(userInput, "\n")
	userInput = strings.TrimSuffix(userInput, "\r")

	if err != nil {
		return "", err
	}

	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput, nil
}
