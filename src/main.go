package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	CAR    = "🚗"
	GOAT   = "🐐"
	CLOSED = "🚪"
)

func secureRandomInt(limit int64) int {
	n, err := rand.Int(rand.Reader, big.NewInt(limit))
	if err != nil {
		panic(err)
	}
	return int(n.Int64())
}

func montyHallSimulation(change bool) bool {
	carPosition := secureRandomInt(3)
	playerChoice := secureRandomInt(3)

	if change {
		for {
			montyChoice := secureRandomInt(3)
			if montyChoice != playerChoice && montyChoice != carPosition {
				playerChoice = 3 - playerChoice - montyChoice
				break
			}
		}
	}

	return playerChoice == carPosition
}

func main() {
	fmt.Println("🎉 Welcome to the Monty Hall Problem Simulator! 🎉")
	fmt.Println("\n📖 Context:")
	fmt.Println("Imagine you're a contestant on a game show. In front of you are three doors: ", CLOSED, CLOSED, CLOSED)
	fmt.Println("Behind one of them is a shiny car ", CAR, "and behind the other two are goats ", GOAT)
	fmt.Println("You pick a door. Monty Hall, the host who knows the secret behind each door, opens another door, always revealing a goat ", GOAT)
	fmt.Println("He then poses a question: 'Do you wish to switch your choice to the other unopened door?'")
	fmt.Println("\n🤔 What's your move? Stick with your initial pick, or switch to the other door?")
	fmt.Println("\nThis simulator will unveil the winning odds for both strategies.")

	tests := []int{10, 100, 1000, 10000, 100000, 1000000}

	for _, testCount := range tests {
		stayWins := 0
		changeWins := 0

		for j := 0; j < testCount; j++ {
			if montyHallSimulation(false) {
				stayWins++
			}
			if montyHallSimulation(true) {
				changeWins++
			}
		}

		stayPercentage := float64(stayWins) * 100 / float64(testCount)
		changePercentage := float64(changeWins) * 100 / float64(testCount)
		percentageDifference := changePercentage - stayPercentage

		fmt.Printf("\n🔍 Results after %d simulations:\n", testCount)
		fmt.Printf("Sticking to the first choice: %d wins (%.2f%%)\n", stayWins, stayPercentage)
		fmt.Printf("Switching doors: %d wins (%.2f%%)\n", changeWins, changePercentage)
		fmt.Printf("🚀 By making the switch, you amplify your winning odds by %.2f%%!\n", percentageDifference)
		fmt.Println("---------------------------------------------------")
	}

	fmt.Println("\n🔢 The Math Unraveled:")
	fmt.Println("At the start, picking the door with the car ", CAR, "has a 1/3 likelihood, leaving a 2/3 chance of selecting a goat ", GOAT)
	fmt.Println("Choosing the car (1/3 chance) and then making a switch ensures a loss.")
	fmt.Println("Picking a goat (2/3 chance) followed by a switch almost guarantees a win with the car behind the other door!")
	fmt.Println("Consequently, by switching, you effectively double your win chances!")
	fmt.Println("\n🧠 Analogy to Ponder:")
	fmt.Println("Visualize the game, but this time with 100 doors. Behind one is a car ", CAR, "and behind the other 99 are goats ", GOAT)
	fmt.Println("You mark your choice. The probability that your door hides the car stands at a slim 1/100.")
	fmt.Println("Monty, being Monty, opens 98 other doors, each revealing a goat. It's now a face-off between your initial choice and one remaining door.")
	fmt.Println("Given the circumstances, would you reconsider your choice and switch? Most would deduce that making the switch in this setting offers a greater winning shot!")
}
