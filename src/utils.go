package main

import (
	"fmt"
	"math/rand"
	"time"
)

// checks if card to be played is valid
func isValid(topCard Card, cardToBePlayed Card) bool {
	return (topCard.Number == cardToBePlayed.Number || topCard.Suit == cardToBePlayed.Suit)
}

// checks if player has any valid card to make a move
func validCardsinHand(cards []Card, cardOnTop Card) bool {
	for i := 0; i < len(cards); i++ {
		if isValid(cardOnTop, cards[i]) {
			return true
		}
	}
	return false
}

// removes the card with the given index
func removeCard(cards []Card, index int) []Card {
	ret := make([]Card, 0)
	ret = append(ret, cards[:index]...)
	return append(ret, cards[index+1:]...)
}

// displays the cards held by the current player
func displayCards(currentPlayer *Player, turn int) {
	fmt.Printf("Current Player: Player-%d\n", turn+1)
	for i := 0; i < len(currentPlayer.Hand); i++ {
		card := currentPlayer.Hand[i]
		fmt.Printf("%d) %d - %s\n", i+1, card.Number, card.Suit)
	}
}

//if cards length is 0 then current player
//is decleared as winner, game ends
func checkWinner(cards []Card) bool {
	return len(cards) <= 0
}


//shuffels the given slice in random order
func shuffel(cards []Card) []Card {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(cards),
		func(i, j int) {
			cards[i], cards[j] = cards[j], cards[i]
		},
	)
	return cards
}


//checks if given card is action card or not
//Action Cards: Ace, Kings, Jacks, Queens
func checkActionCardPlayed(card Card) bool {
	return (card.Number == 1 || card.Number == 11 || card.Number == 12 || card.Number == 13)
}


//draws a card from draw pile and returns the updated drawPile
func drawCardFromPile(currentPlayer *Player, drawPile []Card, cardsToDraw int) []Card {
	var cardsDrawn = make([]Card, 0)
	//draws cards from pile = cardsToDraw
	for i := 0; i < cardsToDraw; i++ {
		cardsDrawn = append(cardsDrawn, drawPile[len(drawPile)-1])
		drawPile = removeCard(drawPile, len(drawPile)-1)
	}
	//updates the current player's hand
	cards := currentPlayer.Hand
	cards = append(cards, cardsDrawn...)
	currentPlayer.updateHand(cards)
	return drawPile
}
