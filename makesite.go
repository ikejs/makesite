package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// get path of dir with posts
	dirName := flag.String("dir", ".", "Path to directory with posts.")
	// get all files from dirName
	files, err := ioutil.ReadDir(*dirName)
	check(err)
	// loop over files
	for _, file := range files {
		// check if file is dir
		if !file.IsDir() && filepath.Ext(file.Name()) == ".txt" {
			// get file path, set contents to fileContents
			fileContents, err := ioutil.ReadFile(file.Name())
			check(err)
			// create new HTML file in system
			newFile, err := os.Create(file.Name() + ".html")
			check(err)
			// create new template
			t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
			// write to HTML file
			err = t.Execute(newFile, string(fileContents))
			check(err)
		}
	}
}
