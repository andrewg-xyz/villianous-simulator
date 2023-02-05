package internal

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
	// Villians
	PrinceJohn = "princejohn"
)

type Villian struct {
	Objective string
	Location  map[string]Location
	Power     int
	Hand      []Card
	Deck      []Card
	FateDeck  []Card
}

type Location struct {
	Actions []Action
}

type Action struct {
	Name string
}

type GameboardEntry struct {
	Villian string
	Type    string
	Name    string
	Action  string
}

type Card struct {
	Villian  string
	Type     string
	Name     string
	Cost     int
	Power    int
	Function string
	Ability  string
}
