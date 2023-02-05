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
	Objective string              `json:"objective"`
	Location  map[string]Location `json:"location"`
	Power     int                 `json:"power"`
	Hand      []Card              `json:"hand"`
	Deck      []Card              `json:"deck"`
	FateDeck  []Card              `json:"fateDeck"`
}

type Location struct {
	Actions []Action `json:"actions"`
}

type Action struct {
	Name string `json:"name"`
}

type GameboardEntry struct {
	Villian string
	Type    string
	Name    string
	Action  string
}

type Card struct {
	Villian  string `json:"villian"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Cost     int    `json:"cost"`
	Power    int    `json:"power"`
	Function string `json:"function"`
	Ability  string `json:"ability"`
}
