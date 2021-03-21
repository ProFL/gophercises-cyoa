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
	ArcTemplate *template.Template
	StoryArcs   *model.StoryArcs
	Content     *embed.FS
}

func (m *HTTPFrontend) Start(initialArc string) {
	mux := defaultMux()
	mux.Handle("/", http.RedirectHandler(fmt.Sprintf("/arcs/%s", initialArc), http.StatusFound))
	mux.Handle("/static/", http.FileServer(http.FS(*m.Content)))

	for arcName, storyArc := range *m.StoryArcs {
		log.Println("Registering handler for", arcName)
		route := fmt.Sprintf("/arcs/%s", arcName)
		mux.Handle(route, &handler.ArcHandler{
			ArcTemplate: m.ArcTemplate,
			StoryArc:    storyArc,
		})
	}

	http.ListenAndServe(":8080", mux)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}
