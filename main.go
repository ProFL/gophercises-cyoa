package main

import (
	"embed"
	"encoding/json"
	"flag"
	htmlTemplate "html/template"
	"log"
	"os"
	textTemplate "text/template"

	"github.com/ProFL/gophercises-cyoa/frontend"
	"github.com/ProFL/gophercises-cyoa/model"
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

func parseStoryArcs(filePath string) (initialArc string, storyArcs map[string]model.StoryArc) {
	var storyArcsJSON []byte
	var err error
	if filePath == "" {
		storyArcsJSON, err = content.ReadFile("gopher.json")
		if err != nil {
			log.Panicln("Failed to load default story arcs file contents", err.Error())
		}
	} else {
		storyArcsJSON, err = os.ReadFile(filePath)
		if err != nil {
			log.Panicln("Failed to load story arcs file contents", err.Error())
		}
	}

	storyArcs = make(map[string]model.StoryArc)
	err = json.Unmarshal(storyArcsJSON, &storyArcs)
	if err != nil {
		log.Panicln("Failed to parse story arcs file contents", err.Error())
	}

	initialArc = "intro"
	_, ok := storyArcs[initialArc]
	if !ok {
		for arcName, storyArc := range storyArcs {
			if storyArc.IsInitialStory {
				initialArc = arcName
				break
			}
		}
		_, ok = storyArcs[initialArc]
		if !ok {
			log.Panicln(
				"No initial story was found in the provided file!",
				"It should have a isInitialStory key with true value",
			)
		}
	}
	log.Println("Initial arc:", initialArc)
	return initialArc, storyArcs
}
