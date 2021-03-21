package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/ProFL/gophercises-cyoa/handlers"
	"github.com/ProFL/gophercises-cyoa/models"
)

//go:embed static/* templates/*
//go:embed gopher.json
var content embed.FS

func main() {
	mux := defaultMux()
	mux.Handle("/", &handlers.IndexHandler{RedirectPath: "/arcs/intro"})
	mux.Handle("/static/", http.FileServer(http.FS(content)))

	arcTemplate, err := template.ParseFS(content, "templates/arc.html")
	if err != nil {
		log.Panicln("Failed to load arc page template", err.Error())
	}
	storyArcs := parseStoryArcs()

	for arcName, storyArc := range storyArcs {
		log.Println("Registering handler for", arcName)
		route := fmt.Sprintf("/arcs/%s", arcName)
		mux.Handle(route, &handlers.ArcHandler{
			ArcTemplate: arcTemplate,
			StoryArc:    storyArc,
		})
	}
	http.ListenAndServe(":8080", mux)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
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
