package frontends

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/ProFL/gophercises-cyoa/handlers"
	"github.com/ProFL/gophercises-cyoa/models"
)

func StartHTTPServer(arcTemplate *template.Template, storyArcs *map[string]models.StoryArc, content *embed.FS) {
	mux := defaultMux()
	mux.Handle("/", &handlers.IndexHandler{RedirectPath: "/arcs/intro"})
	mux.Handle("/static/", http.FileServer(http.FS(*content)))

	for arcName, storyArc := range *storyArcs {
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
