package main;

type Card struct{
	Number int;
	Type string;
}

type Deck struct{
	Deck []Card;
}

type Player struct {
	Name string;
	hand []Card;
}