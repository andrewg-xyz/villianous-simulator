package internal

import (
	"encoding/csv"
	"fmt"
	"os"
)

func CreateGame(villianFilePath string) State {
	records, _ := readData(villianFilePath)

	gameData := convert(records)

	state := constructState()
	for _, dataEntry := range gameData {
		state = updateStateWithEntry(state, dataEntry)
	}
	return state
}

func constructState() State {
	villians := make(map[string]Villian)

	return State{
		Villians: villians,
	}
}

func updateStateWithEntry(state State, dataEntry GameboardEntry) State {
	villian := state.Villians[dataEntry.Villian]
	if len(villian.Location) == 0 {
		location := make(map[string]Location)
		villian.Location = location
	}
	if dataEntry.Type == "space" {
		location := villian.Location[dataEntry.Name]
		location.Actions = append(location.Actions, Action{Name: dataEntry.Action})
		villian.Location[dataEntry.Name] = location
	}
	state.Villians[dataEntry.Villian] = villian
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

type GameboardEntry struct {
	Villian string
	Type    string
	Name    string
	Action  string
}

const (
	// Villian Names
	PrinceJohn = "princejohn"
)

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
