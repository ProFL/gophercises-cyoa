package frontend

import (
	"fmt"
	"log"
	"text/template"

	"github.com/ProFL/gophercises-cyoa/cli"
	"github.com/ProFL/gophercises-cyoa/model"
)

type CLIFrontend struct {
	arcTemplate *template.Template
	story       *model.Story
}

func NewCLIFrontend(arcTemplate *template.Template, story *model.Story) *CLIFrontend {
	return &CLIFrontend{
		arcTemplate: arcTemplate,
		story:       story,
	}
}

func (m *CLIFrontend) Start() {
	arcHandlers := make(map[string]*cli.ArcHandler)
	for arcName, storyArc := range m.story.Arcs {
		log.Println("Loading arc", arcName, "handler")
		arcHandlers[arcName] = &cli.ArcHandler{
			ArcTemplate: m.arcTemplate,
			StoryArc:    storyArc,
		}
	}

	cli.CallClear()
	fmt.Println("Welcome to Create Your Own Adventure - An interactive adventure book")
	fmt.Println()
	fmt.Println("When prompted, type in the number of the desired option to proceed to the next story arc")
	fmt.Println()
	fmt.Println("Press enter to start reading...")
	fmt.Scanln()

	nextStory := m.story.InitialArc
	for nextStory != "" {
		cli.CallClear()
		curStory := nextStory

		var err error
		nextStory, err = arcHandlers[curStory].HandlePresentation()
		if err != nil {
			log.Panicln("Failed to present", curStory, err.Error())
		}
	}
}
