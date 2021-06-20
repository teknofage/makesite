package main

import (
	// "fmt"
	"io/ioutil"
	"os"
	"log"
	"text/template"
	"flag"
	"strings"
	"context"

	translate "cloud.google.com/go/translate/apiv3"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"

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
		}
	}
}

// translateText translates input text and returns translated text.
func translateText(text string) error {
	projectID := "the-ridge-317404"
	sourceLang := "en-US"
	targetLang := "fr"
	// text := "Text you wish to translate"

	ctx := context.Background()
	client, err := translate.NewTranslationClient(ctx)
	if err != nil {
			return fmt.Errorf("NewTranslationClient: %v", err)
	}
	defer client.Close()

	req := &translatepb.TranslateTextRequest{
			Parent:             fmt.Sprintf("projects/%s/locations/global", projectID),
			SourceLanguageCode: sourceLang,
			TargetLanguageCode: targetLang,
			MimeType:           "text/plain", // Mime types: "text/plain", "text/html"
			Contents:           []string{text},
	}

	resp, err := client.TranslateText(ctx, req)
	if err != nil {
			return fmt.Errorf("TranslateText: %v", err)
	}

	// Display the translation for each input text provided
	for _, translation := range resp.GetTranslations() {
			fmt.Fprintf(w, "Translated text: %v\n", translation.GetTranslatedText())
	}

	return nil
}