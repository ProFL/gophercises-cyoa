package frontend

import (
	"fmt"
	"log"
	"text/template"

	"github.com/ProFL/gophercises-cyoa/cli"
	"github.com/ProFL/gophercises-cyoa/cyoa"
)

type CLIFrontend struct {
	arcTemplate *template.Template
	story       *cyoa.Story
}

func NewCLIFrontend(arcTemplate *template.Template, story *cyoa.Story) *CLIFrontend {
	return &CLIFrontend{
		arcTemplate: arcTemplate,
		story:       story,
	}
}

func (m *CLIFrontend) Start() {
	arcHandlers := make(map[string]*cli.ArcHandler)
	for arcName, _ := range m.story.Arcs {
		log.Println("Loading arc", arcName, "handler")
		arc := m.story.Arcs[arcName]
		arcHandlers[arcName] = cli.NewArcHandler(m.arcTemplate, &arc)
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
		nextStory, err = arcHandlers[curStory].Handle()
		if err != nil {
			log.Panicln("Failed to handle", curStory, err.Error())
		}
	}
}
