package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"text/template"
	"flag"
	"strings"
)
// Holds the title and body from the text file: update for stretch challenge
type Blog struct {
	Title string
	Body string
}


func main() {
	// defining flags
	var filename string
	var directory string

	//define the flag, using the pointer to filename variable
	flag.StringVar(&filename, "file", "", "Text file name")
	flag.StringVar(&directory, "dir", "", "Directory Name")
	flag.Parse()
	
	if directory != "" {
		directoryManipulation(directory)
	} else if filename != "" {
		fileManipulation(filename)
	}
}


func fileManipulation(filename string) {
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		// A common use of `panic` is to abort if a function returns an error
        // value that we don’t know how to (or want to) handle. This example
        // panics if we get an unexpected error when creating a new file.
        panic(err)
	}
	// create file and give it the name of the read txt file, 
	// and split the txt filename at the . and give it a html extension
	f, err := os.Create(strings.SplitN(filename, ".", 2)[0] + ".html")
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

func directoryManipulation(directory string) {
	files, err := ioutil.ReadDir(directory) 
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.Name()[len(file.Name())-3:] == "txt" {
			fmt.Println(file.Name())
		}
	}
}