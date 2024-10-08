package interaction

import (
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"
)

type RoundData struct {
	Action           string
	PlayerAttackDmg  int
	PlayerHealValue  int
	MonsterAttackDmg int
	PlayerHealth     int
	MonsterHealth    int
}

func PrintGreeting() {
	myFigure := figure.NewFigure("MONSTER SLAYER", "", true)
	myFigure.Print()
	fmt.Println("Starting a new game...")
	fmt.Println("Good luck!")
}

func ShowAvailableActions(specialAtttackIsAvailable bool) {
	fmt.Println("Please choose your action:")
	fmt.Println("-------------------------")
	fmt.Println("1. Attack Monster")
	fmt.Println("2. Heal Yourself")

	if specialAtttackIsAvailable {
		fmt.Println("3. Special Attack")
	}
}
func PrintRoundStatistics(roundData *RoundData) {
	if roundData.Action == "ATTACK" {
		fmt.Printf("Player attacked the monster for %v damage.\n", roundData.PlayerAttackDmg)
	} else if roundData.Action == "SPECIAL_ATTACK" {
		fmt.Printf("Player performed a special attack the monster for %v damage.\n", roundData.PlayerAttackDmg)
	} else {
		fmt.Printf("Player healed for %v.\n", roundData.PlayerHealValue)
	}
	fmt.Printf("Monster attacked the player for %v damage.\n", roundData.MonsterAttackDmg)
	fmt.Printf("Player health: %v\n", roundData.PlayerHealth)
	fmt.Printf("Monster health: %v\n", roundData.MonsterHealth)
}
func DeclareWinner(winner string) {
	fmt.Println("-------------------------")
	myFigure := figure.NewColorFigure("GAME OVER!", "", "red", true)
	myFigure.Print()
	fmt.Println("-------------------------")
	fmt.Printf("%v won!\n ", winner)
}
func WriteLogFile(rounds *[]RoundData) {
	file, err := os.Create("gamelog.txt")
	if err != nil {
		fmt.Println("Writig to log file failed.")
		return
	}
	for index, value := range *rounds {
		logEntry := map[string]string{
			"Round":                 fmt.Sprint(index + 1),
			"Action":                value.Action,
			"Player Attack Damage":  fmt.Sprint(value.PlayerAttackDmg),
			"Player Heal Value":     fmt.Sprint(value.PlayerHealValue),
			"Monster Attack Damage": fmt.Sprint(value.MonsterAttackDmg),
			"Player Health":         fmt.Sprint(value.PlayerHealth),
			"Monster Health":        fmt.Sprint(value.MonsterHealth),
		}
		logLine := fmt.Sprintln(logEntry)
		_, err := file.WriteString(logLine)
		if err != nil {
			fmt.Println("Writig to log file failed.")
			continue
		}
	}
	file.Close()
	fmt.Println("Wrote data to log")
}
