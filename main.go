package main

import (
	"embed"
	"encoding/json"
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

var execMode string

func init() {
	flag.StringVar(&execMode, "mode", "http", "select execution mode (http/cli)")
	flag.Parse()
}

func main() {
	storyArcs := parseStoryArcs()

	if execMode == "http" {
		arcTemplate, err := htmlTemplate.ParseFS(content, "template/arc.html")
		if err != nil {
			log.Panicln("Failed to load arc page template", err.Error())
		}
		frontend.StartHTTPServer(arcTemplate, &storyArcs, &content)
	} else if execMode == "cli" {
		arcTemplate, err := textTemplate.ParseFS(content, "template/arc.txt")
		if err != nil {
			log.Panicln("Failed to load arc text template", err.Error())
		}
		frontend.StartCLI(arcTemplate, &storyArcs)
	} else {
		log.Panicln(execMode, "is not a valid mode, pick one of [http, cli]")
	}
}

func parseStoryArcs() map[string]model.StoryArc {
	storyArcsJSON, err := content.ReadFile("gopher.json")
	if err != nil {
		log.Panicln("Failed to load story arcs file contents", err.Error())
	}
	storyArcs := make(map[string]model.StoryArc)
	err = json.Unmarshal(storyArcsJSON, &storyArcs)
	if err != nil {
		log.Panicln("Failed to parse story arcs file contents", err.Error())
	}
	return storyArcs
}
