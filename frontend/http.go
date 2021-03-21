package frontend

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/ProFL/gophercises-cyoa/cyoa"
)

type HTTPFrontend struct {
	story    *cyoa.Story
	embedFS  *embed.FS
	template *template.Template
}

func NewHTTPFrontend(template *template.Template, story *cyoa.Story,
	embedFS *embed.FS) *HTTPFrontend {
	return &HTTPFrontend{
		story:    story,
		embedFS:  embedFS,
		template: template,
	}
}

func (m *HTTPFrontend) Start() {
	mux := defaultMux()
	mux.Handle("/", http.RedirectHandler(
		fmt.Sprintf("/arcs/%s", m.story.InitialArc), http.StatusFound))
	mux.Handle("/static/", http.FileServer(http.FS(*m.embedFS)))
	mux.Handle("/arcs/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.EscapedPath()
		pathSegments := strings.Split(path, "/")
		arc := pathSegments[len(pathSegments)-1]
		if arc == "" {
			arc = pathSegments[len(pathSegments)-2]
		}
		if storyArc, ok := m.story.Arcs[arc]; ok {
			log.Println("Serving arc", arc)
			err := m.template.Execute(w, storyArc)
			if err != nil {
				log.Println("Failed to render arc template", err.Error())
				http.Error(w, "Something went wrong when trying to render this arc...",
					http.StatusInternalServerError)
			}
			return
		}
		log.Printf("Arc %s was not found\n", arc)
		http.Error(w, "Arc not found", http.StatusNotFound)
	}))

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}
