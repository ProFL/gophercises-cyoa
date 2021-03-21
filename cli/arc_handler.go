package cli

import (
	"fmt"
	"os"
	"text/template"

	"github.com/ProFL/gophercises-cyoa/model"
)

type ArcHandler struct {
	arcTemplate *template.Template
	storyArc    *model.StoryArc
}

func NewArcHandler(arcTemplate *template.Template, storyArc *model.StoryArc) *ArcHandler {
	return &ArcHandler{
		arcTemplate: arcTemplate,
		storyArc:    storyArc,
	}
}

func (m *ArcHandler) Handle() (string, error) {
	err := m.arcTemplate.Execute(os.Stdout, m.storyArc)
	if err != nil {
		return "", err
	}
	numberOfOptions := len(m.storyArc.Options)
	if numberOfOptions > 0 {
		userInput := -1
		for userInput < 0 || userInput > numberOfOptions {
			fmt.Print("Type in the number of your choice) ")
			fmt.Scanln(&userInput)
			if userInput < 0 || userInput > numberOfOptions {
				fmt.Println("Invalid option, please type a number between 1 and", numberOfOptions)
			}
		}
		return m.storyArc.Options[userInput].Arc, nil
	}
	return "", nil
}
