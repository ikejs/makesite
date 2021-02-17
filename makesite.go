package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"

	libsass "github.com/wellington/go-libsass"
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
		if !file.IsDir() {
			// get file path, set contents to fileContents
			fileContents, err := ioutil.ReadFile(file.Name())
			check(err)
			if filepath.Ext(file.Name()) == ".txt" {
				// create new HTML file in system
				newFile, err := os.Create(file.Name() + ".html")
				check(err)
				// create new template
				t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
				// write to HTML file
				err = t.Execute(newFile, string(fileContents))
				check(err)
			}
			if filepath.Ext(file.Name()) == ".scss" {
				// open sass file
				sassFile, err := os.Open("test.scss")
				if err != nil {
					panic(err)
				}
				defer sassFile.Close()

				// create new css file
				outputFile, err := os.Create("styles.css")
				if err != nil {
					panic(err)
				}
				defer outputFile.Close()

				// instantiate libsass
				cssCompiler, err := libsass.New(outputFile, sassFile)
				if err != nil {
					panic(err)
				}

				// compile css
				if err := cssCompiler.Run(); err != nil {
					panic(err)
				}
			}
		}

	}
}
