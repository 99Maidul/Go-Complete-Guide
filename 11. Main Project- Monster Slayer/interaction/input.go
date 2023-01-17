package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetPlayerChoice(specialAtttackIsAvailable bool) {

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
