package main

import (
	"embed"
	"flag"
	htmlTemplate "html/template"
	"log"
	textTemplate "text/template"

	"github.com/ProFL/gophercises-cyoa/frontend"
)

//go:embed static/* template/*
//go:embed gopher.json
var content embed.FS

func main() {
	var frontendMode, storyFilePath string
	flag.StringVar(&frontendMode, "mode", "http", "select execution mode (http/cli)")
	flag.StringVar(&storyFilePath, "story", "", "path to a custom story JSON file")
	flag.Parse()

	initialArc, storyArcs := parseStoryArcs(storyFilePath)

	var app frontend.Frontend
	if frontendMode == "http" {
		arcTemplate, err := htmlTemplate.ParseFS(content, "template/arc.html")
		if err != nil {
			log.Panicln("Failed to load arc page template", err.Error())
		}
		app = &frontend.HTTPFrontend{
			ArcTemplate: arcTemplate,
			StoryArcs:   &storyArcs,
			Content:     &content,
		}
	} else if frontendMode == "cli" {
		arcTemplate, err := textTemplate.ParseFS(content, "template/arc.txt")
		if err != nil {
			log.Panicln("Failed to load arc text template", err.Error())
		}
		app = &frontend.CLIFrontend{
			ArcTemplate: arcTemplate,
			StoryArcs:   &storyArcs,
		}
	} else {
		log.Panicln(frontendMode, "is not a valid mode, pick one of [http, cli]")
	}

	app.Start(initialArc)
}
