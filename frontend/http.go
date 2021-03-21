package frontend

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/ProFL/gophercises-cyoa/handler"
	"github.com/ProFL/gophercises-cyoa/model"
)

type HTTPFrontend struct {
	arcTemplate *template.Template
	story       *model.Story
	embedFS     *embed.FS
}

func NewHTTPFrontend(arcTemplate *template.Template, story *model.Story,
	embedFS *embed.FS) *HTTPFrontend {
	return &HTTPFrontend{
		arcTemplate: arcTemplate,
		story:       story,
		embedFS:     embedFS,
	}
}

func (m *HTTPFrontend) Start() {
	mux := defaultMux()
	mux.Handle("/", http.RedirectHandler(
		fmt.Sprintf("/arcs/%s", m.story.InitialArc), http.StatusFound))
	mux.Handle("/static/", http.FileServer(http.FS(*m.embedFS)))

	for arcName, storyArc := range m.story.Arcs {
		log.Println("Registering handler for", arcName)
		route := fmt.Sprintf("/arcs/%s", arcName)
		mux.Handle(route, &handler.ArcHandler{
			ArcTemplate: m.arcTemplate,
			StoryArc:    storyArc,
		})
	}

	http.ListenAndServe(":8080", mux)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}
