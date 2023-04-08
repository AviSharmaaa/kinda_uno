package main

import "fmt"

//checks if card played is valid
func isValid(topCard Card, playedCard Card) bool {
	return (topCard.Number == playedCard.Number || topCard.Type == playedCard.Type)
}

//checks if player has any valid card to make a move
func validCardsinHand(cards []Card, cardOnTop Card) bool {
	for i := 0; i < len(cards); i++ {
		if isValid(cardOnTop, cards[i]) {
			return true
		}
	}
	return false
}


//removes the card with the given index
func removeCard(cards []Card, index int) []Card {
	ret := make([]Card, 0)
	ret = append(ret, cards[:index]...)
	return append(ret, cards[index+1:]...)
}


//displays the cards help by the current player
func displayCards(currentPlayer *Player, turn int) {
	fmt.Printf("Current Player: Player%d\n", turn+1)
	for i := 0; i < len(currentPlayer.hand); i++ {
		card := currentPlayer.hand[i]
		fmt.Printf("%d) %d - %s\n", i+1, card.Number, card.Type)
	}
}

func checkWinner(cards []Card) bool {
	return len(cards) <= 0
}
