package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
	"flag"
	"strings"
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
	// defining flag
	var filename string

	//define the flag, using the pointer to filename variable
	flag.StringVar(&filename, "file", "", "Text file name")
	flag.Parse()
	if filename == "" {
		fmt.Println("Wubbalubbadubdub!")
		return 
	}
	
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
