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

var execMode string
var storyFile string

func init() {
	flag.StringVar(&execMode, "mode", "http", "select execution mode (http/cli)")
	flag.StringVar(&storyFile, "story", "", "path to a custom story JSON file")
	flag.Parse()
}

func main() {
	initialArc, storyArcs := parseStoryArcs()

	var front frontend.Frontend
	if execMode == "http" {
		arcTemplate, err := htmlTemplate.ParseFS(content, "template/arc.html")
		if err != nil {
			log.Panicln("Failed to load arc page template", err.Error())
		}
		front = &frontend.HTTPFrontend{
			ArcTemplate: arcTemplate,
			StoryArcs:   &storyArcs,
			Content:     &content,
		}
	} else if execMode == "cli" {
		arcTemplate, err := textTemplate.ParseFS(content, "template/arc.txt")
		if err != nil {
			log.Panicln("Failed to load arc text template", err.Error())
		}
		front = &frontend.CLIFrontend{
			ArcTemplate: arcTemplate,
			StoryArcs:   &storyArcs,
		}
	} else {
		log.Panicln(execMode, "is not a valid mode, pick one of [http, cli]")
	}

	front.Start(initialArc)
}

func parseStoryArcs() (initialArc string, storyArcs map[string]model.StoryArc) {
	var storyArcsJSON []byte
	var err error
	if storyFile == "" {
		storyArcsJSON, err = content.ReadFile("gopher.json")
		if err != nil {
			log.Panicln("Failed to load default story arcs file contents", err.Error())
		}
	} else {
		storyArcsJSON, err = os.ReadFile(storyFile)
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
