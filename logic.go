package main

import (
	"fmt"
)

func getDeck() []Card {
	var cards []Card

	suits := []string{"heart", "spade", "diamond", "club"}

	for _, element := range suits {
		var card Card
		for i := 1; i < 14; i++ {
			card = Card{i, element}
			cards = append(cards, card)
		}
	}

	//shuffels the deck in random order
	cards = shuffel(cards)

	return cards
}

// returns list of players and remianing deck
func getPlayersAndDrawPile(noOfPlayers int) ([]Player, []Card) {
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
	players, drawCards := getPlayersAndDrawPile(noOfPlayers)

	//slice of cards played by players
	discardPile := make([]Card, 0)
	discardPile = append(discardPile, drawCards[len(drawCards)-1])
	drawCards = removeCard(drawCards, len(drawCards)-1)

	playerTurn := 0
	direction := 1

	for {
		//draw game condition
		if len(drawCards) <= 0 {
			fmt.Println("Match Draw, Game Over!!")
			break
		}

		playerTurn %= noOfPlayers
		if playerTurn < 0 {
			playerTurn += noOfPlayers // if playerturn goes to negative integers.
			playerTurn %= noOfPlayers
		}

		//get top card
		cardOnTop := discardPile[len(discardPile)-1]
		fmt.Printf("Card on top:  %d - %s\n", cardOnTop.Number, cardOnTop.Suit)

		//display current players card
		currentPlayer := &players[playerTurn]
		displayCards(currentPlayer, playerTurn)

		//checks if top card is action card
		isActionCard := checkActionCardPlayed(cardOnTop)

		cardPlayed := false

		//checks if player can make a valid move
		if validCardsinHand(currentPlayer.Hand, cardOnTop) {

			//player plays a valid card
			for index, card := range currentPlayer.Hand {
				if isValid(cardOnTop, card) {
					// moves the played car into discard pile
					if isActionCard && card.Number == cardOnTop.Number {
						continue
					}
					fmt.Printf("Player-%d played: %d - %s  \n\n",playerTurn + 1, card.Number, card.Suit)
					cardPlayed = true
					discardPile = append(discardPile, card)

					//updates the hand of current player
					cards := removeCard(currentPlayer.Hand, index)
					currentPlayer.updateHand(cards)
					break
				}
			}

		} else {
			fmt.Printf("No valid cards, drawing from the draw pile!!!\n\n")
			//Draws cards from draw cards pile and updates
			// the current players hand
			drawCards = drawCardFromPile(currentPlayer, drawCards, 1)
		}

		//check if player has cards
		if checkWinner(currentPlayer.Hand) {
			//declares the current player as winner if no cards are left
			//ends the game
			fmt.Printf("Congratulations, Player-%d won, Game Over!!!\n\n", playerTurn+1)
			break
		}

		if isActionCard && cardPlayed {
			//if card was Ace, skip next player's turn
			if cardOnTop.Number == 1 {
				playerTurn += direction
			}

			//if card was king, reverse the direction of game flow
			if cardOnTop.Number == 13 {
				direction *= -1
			}

			//if card was Jack, next player picks 4 cards
			if cardOnTop.Number == 11 {
				drawCards = drawCardFromPile(currentPlayer, drawCards, 4)
			}

			// if card was Queen, next player picks 2 cards
			if cardOnTop.Number == 12 {
				drawCards = drawCardFromPile(currentPlayer, drawCards, 2)
			}
		}
		playerTurn += direction
		fmt.Println("=========================================")
	}
}
