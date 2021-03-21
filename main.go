package main

import (
	"embed"
	"encoding/json"
	"flag"
	"html/template"
	"log"

	"github.com/ProFL/gophercises-cyoa/frontends"
	"github.com/ProFL/gophercises-cyoa/models"
)

//go:embed static/* templates/*
//go:embed gopher.json
var content embed.FS

var execMode string

func init() {
	flag.StringVar(&execMode, "mode", "http", "select execution mode (http/cli)")
	flag.Parse()
}

func main() {
	arcTemplate, err := template.ParseFS(content, "templates/arc.html")
	if err != nil {
		log.Panicln("Failed to load arc page template", err.Error())
	}
	storyArcs := parseStoryArcs()

	if execMode == "http" {
		frontends.StartHTTPServer(arcTemplate, &storyArcs, &content)
	} else if execMode == "cli" {
		log.Panicln("cli mode is not yet implemented")
	} else {
		log.Panicln(execMode, "is not a valid mode, pick one of [http, cli]")
	}
}

func parseStoryArcs() map[string]models.StoryArc {
	storyArcsJSON, err := content.ReadFile("gopher.json")
	if err != nil {
		log.Panicln("Failed to load story arcs file contents", err.Error())
	}
	storyArcs := make(map[string]models.StoryArc)
	err = json.Unmarshal(storyArcsJSON, &storyArcs)
	if err != nil {
		log.Panicln("Failed to parse story arcs file contents", err.Error())
	}
	return storyArcs
}
