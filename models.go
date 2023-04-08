package main

type Card struct {
	Number int
	Suit   string
}

type Player struct {
	Name string
	Hand []Card
}

//model methods
func (p *Player) updateHand(updatedHand []Card) {
	p.Hand = updatedHand
}
