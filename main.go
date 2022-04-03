package main

import (
	"github.com/EviCoach/monsterslayer/actions"
	"github.com/EviCoach/monsterslayer/interactions"
)

var currentRound int = 0

func main() {
	startGame()

	winner := "" // "Player" || "Monster" || ""

	for winner == "" {
		winner = executeRound()
	}

	endGame(winner)
}

func startGame() {
	interactions.PrintGreeting()
}
func executeRound() string {
	currentRound = currentRound + 1
	isSpecialRound := currentRound%3 == 0

	interactions.ShowAvailableActions(isSpecialRound)
	userChoice := interactions.GetPlayerChoice(isSpecialRound)

	var playerAttackDmg int
	var playerHealValue int
	var monsterAttackDmg int

	if userChoice == "ATTACK" {
		playerAttackDmg = actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		playerHealValue = actions.HealPlayer()
	} else {
		playerAttackDmg = actions.AttackMonster(true)
	}

	monsterAttackDmg = actions.AttackPlayer()
	playerHealth, monsterHealth := actions.GetHealthAmounts()

	roundData := interactions.RoundData{
		Action: userChoice,
		PlayerHealth: playerHealth,
		MonsterHealth: monsterHealth,
		PlayerAttackDmg: playerAttackDmg,
		PlayerHealValue: playerHealValue,
		MonsterAttackDmg: monsterAttackDmg,
	}

	interactions.PrintRoundStatistics(&roundData)

	if playerHealth <= 0 {
		return "Monster"
	} else if monsterHealth <= 0 {
		return "Player"
	}

	return ""
}
func endGame(winner string) {
	interactions.DeclareWinner(winner)
}
