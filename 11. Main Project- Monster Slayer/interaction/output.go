package interaction

import "fmt"

func PrintGreeting() {
	fmt.Println("MONSTER SLAYER")
	fmt.Println("Starting a new game...")
	fmt.Println("Good luck!")
}

func ShowAvailableActions(specialAtttackIsAvailable bool) {
	fmt.Println("Please choose your action:")
	fmt.Println("-------------------------")
	fmt.Println("1. Attack Monster")
	fmt.Println("2. Heal Yourself")

	if specialAtttackIsAvailable == true {
		fmt.Println("3. Special Attack")
	}
}
