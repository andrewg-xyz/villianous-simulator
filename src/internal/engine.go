package internal

import (
	"math/rand"
)

func ShuffleNTimes(cards []Card, n int) []Card {
	for i := 0; i < n; i++ {
		cards = shuffle(cards)
	}
	return cards
}

func shuffle(cards []Card) []Card {
	rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
	return cards
}
