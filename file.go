package main

import (
	"bytes"
	"html/template"
	"io/fs"
	"log"
	"os"
	"os/exec"
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

		// fmt.Println(templates, rendered.Done, difference)

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
	fileName := strings.Split(path, "/")[1]
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

	os.MkdirAll("./dist/"+fileName, os.ModePerm)

	fileErr := os.WriteFile("./dist/"+fileName+"/index.html", distContent.Bytes(), 0644)
	errorCheck(fileErr)
}

func compileTypescript(path string) {
	// fileName := strings.Replace(path, "src/", "", 1)
	fileName := strings.Split(path, "/")[1]

	command := "npx swc src/" + fileName + "/ts/*.ts -o dist/" + fileName + "/js/index.js"
	parts := strings.Fields(command)
	println("Compiling Typescript: ", "src/"+fileName+"/ts/*.ts")

	out, err := exec.Command(parts[0], parts[1:]...).Output()

	if false {
		println("output: " + string(out))
	}

	if err != nil {
		log.Fatal(err)
	}
}

func minifyCSS(path string) {
	// fileName := strings.Replace(path, "src/", "", 1)
	fileName := strings.Split(path, "/")[1]

	command := "npx postcss src/" + fileName + "/css/style.css -o dist/" + fileName + "/css/style.css"
	parts := strings.Fields(command)
	println("Minifying CSS: ", "src/"+fileName+"/css/style.css")

	out, err := exec.Command(parts[0], parts[1:]...).Output()

	if false {
		println("output: " + string(out))
	}

	if err != nil {
		log.Fatal(err)
	}

}
