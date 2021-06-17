package main

import (
	// "fmt"
	"io/ioutil"
	"os"
	"text/template"
)
type entry struct {
	Name string
	Done bool
}

type ToDo struct {
	User string
	List []entry
}


func main() {
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		// A common use of `panic` is to abort if a function returns an error
        // value that we don’t know how to (or want to) handle. This example
        // panics if we get an unexpected error when creating a new file.
        panic(err)
	}

	f, err := os.Create("first-post.html")
    if err != nil {
        panic(err)
    }

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
        err = t.Execute(f, string(fileContents))
        if err != nil {
          panic(err)
        }
		f.Close()
}
