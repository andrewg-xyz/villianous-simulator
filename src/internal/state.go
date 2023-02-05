package internal

type State struct {
	Villians map[string]Villian
}

type Villian struct {
	Objective string
	Location  map[string]Location
	Power int
	Hand []Card
	Deck []Card
	FateDeck []Card
}

type Location struct {
	Actions []Action
}

type Action struct {
	Name string
}

const (
	// Actions
	GainPower      = "Gain Power"
	PlayCard       = "Play Card"
	Activate       = "Activate"
	Fate           = "Fate"
	MoveItemOrAlly = "Move Item or Ally"
	MoveHero       = "Move Hero"
	Vanquish       = "Vanquish"
	Discard        = "Discard"
)
