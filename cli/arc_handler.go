package cli

import (
	"fmt"
	"os"
	"text/template"

	"github.com/ProFL/gophercises-cyoa/model"
)

type ArcHandler struct {
	ArcTemplate *template.Template
	StoryArc    model.StoryArc
}

func (m *ArcHandler) HandlePresentation() (string, error) {
	err := m.ArcTemplate.Execute(os.Stdout, m.StoryArc)
	if err != nil {
		return "", err
	}
	numberOfOptions := len(m.StoryArc.Options)
	if numberOfOptions > 0 {
		userInput := -1
		for userInput < 0 || userInput > numberOfOptions {
			fmt.Print("Type in the number of your choice) ")
			fmt.Scanln(&userInput)
			if userInput < 0 || userInput > numberOfOptions {
				fmt.Println("Invalid option, please type a number between 1 and", numberOfOptions)
			}
		}
		return m.StoryArc.Options[userInput].Arc, nil
	}
	return "", nil
}
