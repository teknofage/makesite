package main

import (
	// "fmt"
	"io/ioutil"
	"os"
	"log"
	"text/template"
	"flag"
	"strings"
	"time"

	"github.com/hajimehoshi/oto"
	"github.com/tosone/minimp3"
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
			fileManipulation(file.Name())
		} else if file.Name()[len(file.Name())-3:] == "mp3"{
			decodeMp3(file.Name())
		}
	} 
}

func decodeMp3(mp3name string) {
	var err error

	var file []byte
	if file, err = ioutil.ReadFile(mp3name); err != nil {
		log.Fatal(err)
	}

	var dec *minimp3.Decoder
	var data []byte
	if dec, data, err = minimp3.DecodeFull(file); err != nil {
		log.Fatal(err)
	}

	var context *oto.Context
	if context, err = oto.NewContext(dec.SampleRate, dec.Channels, 2, 1024); err != nil {
		log.Fatal(err)
	}

	var player = context.NewPlayer()
	player.Write(data)

	<-time.After(time.Second)

	dec.Close()
	if err = player.Close(); err != nil {
		log.Fatal(err)
	}
}
