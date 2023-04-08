package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter the number of players")

	var noOfPlayers int

	//gets the no of players
	for {
		//if the number of players > 4 or players < 2
		//prompts to make a correct input
		fmt.Scanln(&noOfPlayers)
		if noOfPlayers > 1 && noOfPlayers < 5 {
			break
		}
		if noOfPlayers < 2 {
			fmt.Println("Players cannot be less than 2")
		} else if noOfPlayers > 4 {
			fmt.Println("Players cannot be greater than 4")
		}
	}

	//starts actual game
	playGame(noOfPlayers)

}
