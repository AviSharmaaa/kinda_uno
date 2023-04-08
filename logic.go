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
