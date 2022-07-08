package main

import (
	"bytes"
	"html/template"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func getTemplates(delta bool) []string {
	var templates []string
	templates = append(templates, findFileByExtension("./src", ".html")...)

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
	writeColor("Rending template: ", "red", false)
	parts := strings.Split(path, "/")
	moduleName := strings.Split(parts[2], ".")
	println(moduleName[0])
	fileName := parts[1]
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
	errorCheck(templateErr)

	os.MkdirAll("./dist/"+fileName, os.ModePerm)

	// for i := 1; i <= 50000; i++ {
	// 	index := strconv.Itoa(i)
	// 	fileErr := os.WriteFile("./dist/"+fileName+"/index-"+index+".html", distContent.Bytes(), 0644)
	// 	errorCheck(fileErr)
	// }

	fileErr := os.WriteFile("./dist/"+fileName+"/index.html", distContent.Bytes(), 0644)
	errorCheck(fileErr)
}

func compileTypescript() {
	command := "npx swc src -d dist -s"
	commandParts := strings.Fields(command)
	writeColor("Compiling Typescript..", "blue", true)

	out, err := exec.Command(commandParts[0], commandParts[1:]...).Output()

	if false {
		println("output: " + string(out))
	}

	errorCheck(err)
}

func minifyCSS() {
	command := "npx postcss src/**/*.css --base src --dir dist"
	commandParts := strings.Fields(command)
	writeColor("Minifying CSS..", "green", true)

	out, err := exec.Command(commandParts[0], commandParts[1:]...).Output()

	if false {
		println("output: " + string(out))
	}

	errorCheck(err)

}
