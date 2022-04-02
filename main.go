package main

import (
	"fmt"

	"github.com/EviCoach/monsterslayer/interactions"
)

var currentRound int = 0

func main() {
	startGame()

	winner := "" // "Player" || "Monster" || ""

	for winner == "" {
		winner = executeRound()
	}

	endGame()
}

func startGame() {
	interactions.PrintGreeting()
}
func executeRound() string {
	currentRound = currentRound + 1
	isSpecialRound := currentRound%3 == 0

	interactions.ShowAvailableActions(isSpecialRound)
	userChoice := interactions.GetPlayerChoice(isSpecialRound)
	
	if userChoice == "ATTACK"{

	} else if userChoice == "HEAL" {

	}
	return ""
}
func endGame() {}
