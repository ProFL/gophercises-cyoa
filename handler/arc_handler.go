package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/ProFL/gophercises-cyoa/model"
)

type ArcHandler struct {
	ArcTemplate *template.Template
	StoryArc    model.StoryArc
}

func (m *ArcHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	err := m.ArcTemplate.Execute(res, m.StoryArc)
	if err != nil {
		log.Println("Failed to render arc template", err.Error())
		res.WriteHeader(http.StatusInternalServerError)
	}
}