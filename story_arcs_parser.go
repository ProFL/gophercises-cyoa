package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/ProFL/gophercises-cyoa/model"
)

func parseStoryArcs(filePath string) (initialArc string, storyArcs map[string]model.StoryArc) {
	storyArcsJSON, err := extractJSONContents(filePath)
	if err != nil {
		log.Panicln(storyArcsJSON, err.Error())
	}

	storyArcs = make(map[string]model.StoryArc)
	err = json.Unmarshal(storyArcsJSON, &storyArcs)
	if err != nil {
		log.Panicln("Failed to parse story arcs file contents", err.Error())
	}

	initialArc, err = extractInitialArc(&storyArcs)
	if err != nil {
		log.Panicln(err.Error())
	}
	log.Println("Initial arc:", initialArc)
	return initialArc, storyArcs
}

func extractJSONContents(filePath string) ([]byte, error) {
	var storyArcsJSON []byte
	var err error
	if filePath == "" {
		storyArcsJSON, err = content.ReadFile("gopher.json")
		if err != nil {
			return []byte("Failed to load default story arcs file contents"), err
		}
	} else {
		storyArcsJSON, err = os.ReadFile(filePath)
		if err != nil {
			return []byte("Failed to load story arcs file contents"), err
		}
	}
	return storyArcsJSON, nil
}

func extractInitialArc(storyArcs *map[string]model.StoryArc) (string, error) {
	initialArc := "intro"
	_, ok := (*storyArcs)[initialArc]
	if !ok {
		for arcName, storyArc := range *storyArcs {
			if storyArc.IsInitialStory {
				initialArc = arcName
				break
			}
		}
		_, ok = (*storyArcs)[initialArc]
		if !ok {
			return "", errors.New(
				"no initial story was found in the provided file, " +
					`it should have a "isInitialStory" key with true value`,
			)
		}
	}
	return initialArc, nil
}
