package main

import (
	"github.com/99Maidul/monsterslayer/interaction"
)

var currentRound = 0

func main() {
	startGame()

	winner := ""

	for winner == "" {
		winner = executeRound()
	}

	endGame()
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0

	interaction.ShowAvailableActions(isSpecialRound)
}

func endGame() {
}
