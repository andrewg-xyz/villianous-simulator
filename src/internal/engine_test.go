package internal

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShuffle(t *testing.T){
	locationFilePath, _ := filepath.Abs("../../data/pj_location.csv")
	deckFilePath, _ := filepath.Abs("../../data/pj_deck.csv")

	state := CreateGame(locationFilePath, deckFilePath)
	deck := state.Villians["princejohn"].Deck 
	fourthCard := deck[3].Name
	newDeck := ShuffleNTimes(deck,7)

	// check that (random) 5th card has moved
	assert.NotEqual(t, fourthCard, newDeck[3].Name)
}