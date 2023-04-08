package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getDeck() []Card {
	var cards []Card

	suits := []string{"heart", "spade", "diamond", "club"}

	for _, element := range suits {
		for i := 1; i < 14; i++ {
			card := Card{i, element}

			cards = append(cards, card)
		}
	}

	//shuffels the deck in random order
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards),
		func(i, j int) {
			cards[i], cards[j] = cards[j], cards[i]
		},
	)

	return cards
}

// returns list of players and remianing deck
func getPlayers(noOfPlayers int) ([]Player, []Card) {
	//get the cards deck
	cards := getDeck()

	var players []Player

	for i := 1; i < noOfPlayers+1; i++ {
		playerName := fmt.Sprintf("Player %d", i)
		var hand []Card

		for j := 0; j < 5; j++ {
			hand = append(hand, cards[j])
			cards = append(cards[:j], cards[j+1:]...)
		}
		player := Player{playerName, hand}

		players = append(players, player)
	}

	return players, cards
}

func playGame(noOfPlayers int) {

	//gets players slice, drawCards slice
	players, drawCards := getPlayers(noOfPlayers)

	//slice of cards played by players
	discardPile := []Card{drawCards[0]}
	drawCards = removeCard(drawCards, 0)

	playerTurn := 0

	for {
		//draw game condition
		if len(drawCards) <= 0 {
			fmt.Println("Game Over!!")
			break
		}

		playerTurn %= noOfPlayers
		if playerTurn < 0 {
			playerTurn += noOfPlayers // if playerturn goes to negative integers.
			playerTurn %= noOfPlayers
		}

		//get top card
		cardOnTop := discardPile[len(discardPile)-1]
		fmt.Printf("Card on top:  %d - %s\n", cardOnTop.Number, cardOnTop.Type)

		//display current players card
		currentPlayer := &players[playerTurn]
		displayCards(currentPlayer, playerTurn)

		//checks if player can make a valid move
		if validCardsinHand(currentPlayer.hand, cardOnTop) {
			fmt.Println("Choose a card")

			//player plays a valid card
			for {
				var card int
				fmt.Scan(&card)
				fmt.Println(currentPlayer.hand[card-1])
				if isValid(cardOnTop, currentPlayer.hand[card-1]) {

					//moves the played car into discard pile
					discardPile = append(discardPile, currentPlayer.hand[card-1])

					//updates the hand of current player
					cards := removeCard(currentPlayer.hand, card-1)
					currentPlayer.updateHand(cards)
					break
				} else {
					fmt.Println("Choose a valid card")
				}
			}

			//check if player has cards
			if checkWinner(currentPlayer.hand) {
				//declares the current player as winner if no cards are left
				//ends the game
				fmt.Printf("Congratulations, Player-%d won, Game Over!!!\n\n", playerTurn + 1)
				break
			}
		} else {
			fmt.Printf("No valid cards, drawing from the draw pile!!!\n\n")
			//updates the current player's hand
			cards := currentPlayer.hand
			cards = append(cards, drawCards[len(drawCards)-1])
			currentPlayer.updateHand(cards)

			//removes the card picked from draw pile
			drawCards = removeCard(drawCards, len(drawCards)-1)
		}

		playerTurn += 1
	}
}
