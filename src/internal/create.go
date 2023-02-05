package internal

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func CreateGame(villianFilePath string, cardsFilePath string) State {
	records, _ := readData(villianFilePath)
	cards, _ := readData(cardsFilePath)

	gameData := convert(records)
	cardsData := convertCards(cards)

	state := constructState()
	state = constructRealm(state, gameData)
	state = contructDecks(state, cardsData)

	return state
}

func contructDecks(state State, cardsData []Card) State {
	for _, card := range cardsData {
		villian := state.Villians[card.Villian]
		if card.Type == "villian card" {
			villian.Deck = append(villian.Deck, card)
		} else if card.Type == "fate card" {
			villian.FateDeck = append(villian.FateDeck, card)
		}
		state.Villians[card.Villian] = villian
	}
	return state
}

func constructState() State {
	villians := make(map[string]Villian)

	return State{
		Villians: villians,
	}
}

func constructRealm(state State, dataEntries []GameboardEntry) State {
	for _, dataEntry := range dataEntries {
		villian := state.Villians[dataEntry.Villian]
		if len(villian.Location) == 0 {
			location := make(map[string]Location)
			villian.Location = location
		}
		if dataEntry.Type == "space" {
			location := villian.Location[dataEntry.Name]
			location.Actions = append(location.Actions, Action{Name: dataEntry.Action})
			villian.Location[dataEntry.Name] = location
		} else if dataEntry.Type == "objective" {
			villian.Objective = dataEntry.Action
		}
		state.Villians[dataEntry.Villian] = villian
	}
	return state
}

func convert(records [][]string) []GameboardEntry {
	data := []GameboardEntry{}
	for _, record := range records {
		data = append(data, GameboardEntry{
			Villian: record[0],
			Type:    record[1],
			Name:    record[2],
			Action:  record[3],
		})
	}
	return data
}

func convertCards(records [][]string) []Card {
	cards := []Card{}
	for _, record := range records {
		cost, _ := strconv.Atoi(record[3])
		power, _ := strconv.Atoi(record[4])
		cards = append(cards, Card{
			Villian:  record[0],
			Type:     record[1],
			Name:     record[2],
			Cost:     cost,
			Power:    power,
			Function: record[5],
			Ability:  record[6],
		})
	}
	return cards
}

func readData(fileName string) ([][]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return [][]string{}, err
	}
	defer f.Close()
	r := csv.NewReader(f)
	// skip first line
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}
	records, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return records, nil
}
