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
	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	err = t.Execute(os.Stdout, string(fileContents))
	check(err)
}
