package main

//entry point for the game

import (
	"fmt"
)

func main() {
	fmt.Println("Enter the number of players")

	var noOfPlayers int

	//gets the no of players
	//make this a function
	for ok := true; ok; ok = (noOfPlayers < 2 || noOfPlayers > 4) {
		fmt.Scanln(&noOfPlayers)
		if noOfPlayers < 2 {
			fmt.Println("Players cannot be less than 2")
		} else if noOfPlayers > 4 {
			fmt.Println("Players cannot be greater than 4")
		}
	}

	//starts actual game
	playGame(noOfPlayers)

}
