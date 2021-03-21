package frontend

import (
	"fmt"
	"log"
	"text/template"

	"github.com/ProFL/gophercises-cyoa/cli"
	"github.com/ProFL/gophercises-cyoa/model"
)

func StartCLI(arcTemplate *template.Template, storyArcs *map[string]model.StoryArc) {
	arcHandlers := make(map[string]*cli.ArcHandler)
	for arcName, storyArc := range *storyArcs {
		log.Println("Loading arc", arcName, "handler")
		arcHandlers[arcName] = &cli.ArcHandler{
			ArcTemplate: arcTemplate,
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

	nextStory := "intro"
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
