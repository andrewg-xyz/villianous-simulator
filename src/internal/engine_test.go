package internal

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShuffle(t *testing.T) {
	locationFilePath, _ := filepath.Abs("../../data/pj_location.csv")
	deckFilePath, _ := filepath.Abs("../../data/pj_deck.csv")

	state := CreateGame(locationFilePath, deckFilePath)
	deck := state.Villians["princejohn"].Deck
	fourthCard := deck[3].Name
	newDeck := ShuffleNTimes(deck, 7)

	// check that (random) 5th card has moved
	assert.NotEqual(t, fourthCard, newDeck[3].Name)
}

func TestTakeTurn(t *testing.T) {
	// WIP Mocking
	locationFilePath, _ := filepath.Abs("../../data/pj_location.csv")
	deckFilePath, _ := filepath.Abs("../../data/pj_deck.csv")

	state := CreateGame(locationFilePath, deckFilePath)

	for _, action := range state.Villians["princejohn"].Location["sherwood forest"].Actions {
		state = TakeAction(state, action)
	}
	for _, action := range state.Villians["princejohn"].Location["friar tucks church"].Actions {
		state = TakeAction(state, action)
	}
	for _, action := range state.Villians["princejohn"].Location["nottingham"].Actions {
		state = TakeAction(state, action)
	}
	for _, action := range state.Villians["princejohn"].Location["the jail"].Actions {
		state = TakeAction(state, action)
	}
}