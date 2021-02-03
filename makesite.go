package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// get path of file
	fileName := flag.String("file", "latest-post.txt", "Path of the .txt file.")
	// parse input text
	flag.Parse()
	// read file at fileName path, set contents to fileContents
	fileContents, err := ioutil.ReadFile(*fileName)
	check(err)
	// create new HTML file in system
	newFile, err := os.Create("first-post.html")
	check(err)
	// create new template
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	// write to HTML file
	err = t.Execute(newFile, string(fileContents))
	check(err)
}
