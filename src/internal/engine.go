package internal

import (
	"math/rand"
	"regexp"
	"strconv"

	"github.com/rs/zerolog/log"
)

const (
	gainPowerPattern = `gain (\d+) power`
)

var gainPowerRegex = regexp.MustCompile(gainPowerPattern)

// actionionator
func TakeAction(state State, action Action) State {
	log.Debug().Msgf("Taking action: %s", action.Name)

	switch action.Name {
	case gainPowerRegex.FindString(action.Name):
		matches := gainPowerRegex.FindStringSubmatch(action.Name)
		x, _ := strconv.Atoi(matches[1])
		power := state.Villians["princejohn"].Power
		power += x
		villian := state.Villians["princejohn"]
		villian.Power = power
		state.Villians["princejohn"] = villian
	default:
		log.Info().Msgf("Action not found: %s", action.Name)
	}
	return state
}

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
