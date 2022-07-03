package main

import (
	"strings"
)

func main() {
	// Find or create history.json
	createHistory(History{})

	// Find files with .html extension
	pages := getTemplates(true)

	// Render templates
	for i := 0; i < len(pages); i++ {
		println("--" + strings.Split(pages[i], "/")[1] + "-----------------------------")
		go renderTemplate(pages[i])
		go minifyCSS(pages[i])
		compileTypescript(pages[i])
		println("")
	}

	// Update history.json
	updateHistory()
}
