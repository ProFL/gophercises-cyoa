package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/ProFL/gophercises-cyoa/cyoa"
)

type ArcHandler struct {
	arcTemplate *template.Template
	storyArc    *cyoa.StoryArc
}

func NewArcHandler(arcTemplate *template.Template, storyArc *cyoa.StoryArc) *ArcHandler {
	return &ArcHandler{
		arcTemplate: arcTemplate,
		storyArc:    storyArc,
	}
}

func (m *ArcHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	err := m.arcTemplate.Execute(res, *m.storyArc)
	if err != nil {
		log.Println("Failed to render arc template", err.Error())
		res.WriteHeader(http.StatusInternalServerError)
	}
}
