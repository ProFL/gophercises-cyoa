package frontend

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/ProFL/gophercises-cyoa/handler"
	"github.com/ProFL/gophercises-cyoa/cyoa"
)

type HTTPFrontend struct {
	arcTemplate *template.Template
	story       *cyoa.Story
	embedFS     *embed.FS
}

func NewHTTPFrontend(arcTemplate *template.Template, story *cyoa.Story,
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

	for arcName, _ := range m.story.Arcs {
		log.Println("Registering handler for", arcName)
		arc := m.story.Arcs[arcName]
		mux.Handle(fmt.Sprintf("/arcs/%s", arcName), handler.NewArcHandler(m.arcTemplate, &arc))
	}

	http.ListenAndServe(":8080", mux)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}
