package main

import (
	"encoding/json"
	"io/ioutil"
)

type History struct {
	Done []string `json:"Done"`
}

func createHistory(update History) {
	var history History

	if update.Done != nil {
		history = update
	} else {
		history = History{
			Done: []string{},
		}
	}

	file, _ := json.MarshalIndent(history, "", "	")

	_ = ioutil.WriteFile("history.json", file, 0644)

	// println("history.json saved successfully!")
	return

}

func readHistory() History {
	content, fileErr := ioutil.ReadFile("history.json")
	errorCheck(fileErr)

	history := History{}

	jsonErr := json.Unmarshal(content, &history)
	errorCheck(jsonErr)

	// fmt.Println("history:", history.Done)

	return history
}

func updateHistory() {
	// println("Updating history.")
	history := readHistory()
	templates := getTemplates(false)

	difference := difference(history.Done, templates)

	history.Done = append(history.Done, difference...)

	createHistory(history)
}
