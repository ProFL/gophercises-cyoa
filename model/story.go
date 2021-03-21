package model

import (
	"embed"
	"encoding/json"
	"io"
	"log"
	"os"
)

type StoryArcOption struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type StoryArc struct {
	Title      string           `json:"title"`
	Paragraphs []string         `json:"story"`
	Options    []StoryArcOption `json:"options"`
	IsStart    bool             `default:"false" json:"isStart"`
}

type StoryArcs map[string]StoryArc

type Story struct {
	InitialArc string
	Arcs       StoryArcs
}

func ParseStory(filePath string, fs *embed.FS) *Story {
	arcs := loadArcs(filePath, fs)
	initialArc := extractInitialArc(&arcs)
	return &Story{
		InitialArc: initialArc,
		Arcs:       arcs,
	}
}

func loadArcs(filePath string, fs *embed.FS) (storyArcs StoryArcs) {
	var file io.ReadCloser
	var err error
	if fs != nil {
		file, err = fs.Open(filePath)
		if err != nil {
			log.Fatalln("failed to open embedded story arcs file", err.Error())
		}
	} else {
		file, err = os.Open(filePath)
		if err != nil {
			log.Fatalln("failed to open external story arcs file", err.Error())
		}
	}
	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			log.Println("failed to close story arcs file", closeErr.Error())
		}
	}()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&storyArcs); err != nil {
		log.Fatalln("failed to decode story arcs file contents", err.Error())
	}
	return storyArcs
}

func extractInitialArc(storyArcs *StoryArcs) string {
	initialArc := "intro"
	_, ok := (*storyArcs)[initialArc]
	if !ok {
		for arcName, storyArc := range *storyArcs {
			if storyArc.IsStart {
				initialArc = arcName
				break
			}
		}
		_, ok = (*storyArcs)[initialArc]
		if !ok {
			log.Fatalln("no initial arc found, expected an arc to be named intro or have isStart true")
		}
	}
	return initialArc
}
