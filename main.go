package main

import (
	"embed"
	"flag"
	htmlTemplate "html/template"
	"log"
	textTemplate "text/template"

	"github.com/ProFL/gophercises-cyoa/frontend"
	"github.com/ProFL/gophercises-cyoa/model"
)

//go:embed static/* template/*
//go:embed gopher.json
var content embed.FS

func main() {
	frontendMode := flag.String("mode", "http", "select execution mode (http/cli)")
	storyFilePath := flag.String("story", "", "path to a custom story JSON file")
	flag.Parse()

	var story *model.Story
	if *storyFilePath != "" {
		story = model.ParseStory(*storyFilePath, nil)
	} else {
		story = model.ParseStory("gopher.json", &content)
	}

	var app frontend.Frontend
	if *frontendMode == "http" {
		arcTemplate, err := htmlTemplate.ParseFS(content, "template/arc.html")
		if err != nil {
			log.Panicln("Failed to load arc page template", err.Error())
		}
		app = frontend.NewHTTPFrontend(arcTemplate, story, &content)
	} else if *frontendMode == "cli" {
		arcTemplate, err := textTemplate.ParseFS(content, "template/arc.txt")
		if err != nil {
			log.Panicln("Failed to load arc text template", err.Error())
		}
		app = frontend.NewCLIFrontend(arcTemplate, story)
	} else {
		log.Panicln(frontendMode, "is not a valid mode, pick one of [http, cli]")
	}

	app.Start()
}
