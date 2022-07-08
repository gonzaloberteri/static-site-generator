package main

func main() {
	// Find or create history.json
	createHistory(History{})

	// Find files with .html extension
	pages := getTemplates(true)

	compileTypescript()
	minifyCSS()

	// Render templates
	for i := 0; i < len(pages); i++ {
		go renderTemplate(pages[i])
	}

	// Update history.json
	updateHistory()
}
