package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func getTemplates(delta bool) []string {
	var templates []string
	for _, page := range findFileByExtension("./src", ".html") {
		templates = append(templates, page)
	}

	if delta {
		rendered := readHistory()
		difference := difference(templates, rendered.Done)

		fmt.Println(templates, rendered.Done, difference)

		return difference
	}

	return templates
}

func findFileByExtension(root, ext string) []string {
	var a []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, err error) error {
		errorCheck(err)
		if filepath.Ext(d.Name()) == ext {
			a = append(a, s)
		}
		return nil
	})
	return a
}

func readFile(path string) []byte {
	dat, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return dat
}

func renderTemplate(path string) {
	println("Rending template: ", path)
	fileName := strings.Replace(path, "src/", "", 1)
	// srcContent := readFile(path)
	distContent := new(bytes.Buffer)

	t, err := template.ParseFiles(path)
	errorCheck(err)

	type Post struct {
		Id        int
		Title     string
		CreatedOn time.Time
		Copies    []int
	}

	post := &Post{Id: 19, Title: "Test", Copies: []int{1, 2, 3, 4, 5, 6, 7}}

	templateErr := t.Execute(distContent, post)
	//fmt.Println(distContent)
	errorCheck(templateErr)

	fileErr := os.WriteFile("./dist/"+fileName, distContent.Bytes(), 0644)
	errorCheck(fileErr)
}
