package main

func main() {
	// Find files with .html extension
	var pages []string
	for _, page := range find("./src", ".html") {
		pages = append(pages, page)
	}

	// Render templates
	for i := 0; i < len(pages); i++ {
		createFile(pages[i])
	}
}
