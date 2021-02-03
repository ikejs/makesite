package main

import (
	"io/ioutil"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fileContents, err := ioutil.ReadFile("first-post.txt")
	check(err)
	newFile, err := os.Create("first-post.html")
	check(err)
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	err = t.Execute(newFile, string(fileContents))
	check(err)
}
