package main

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func find(root, ext string) []string {
	var a []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
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

func createFile(path string) {
	fileName := strings.Replace(path, "src/", "", 1)
	content := readFile(path)

	file := os.WriteFile("./dist/"+fileName, content, 0644)
	if file != nil {
		log.Fatal(file)
	}
}

func main() {
	var pages []string
	for _, page := range find("./src", ".html") {
		pages = append(pages, page)
	}

	//fmt.Printf("len=%d cap=%d %v\n", len(pages), cap(pages), pages)

	for i := 0; i < len(pages); i++ {
		createFile(pages[i])
	}
}
