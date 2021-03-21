package main

import (
	"encoding/json"
	"testing"

	"github.com/ProFL/gophercises-cyoa/model"
)

func Test_parseStoryArcs(t *testing.T) {
	expectedInitialArc := "intro"
	initialArc, storyArcs := parseStoryArcs("")
	if initialArc != expectedInitialArc {
		t.Fatalf("Embedded initial arc expected to be %s, but was: %s", expectedInitialArc, initialArc)
	}
	_, ok := storyArcs[initialArc]
	if !ok {
		t.Fatalf("Embedded initial arc (%s) expected to be an available key", initialArc)
	}

	expectedInitialArc = "main"
	initialArc, storyArcs = parseStoryArcs("gopher-no-intro.json")
	if initialArc != expectedInitialArc {
		t.Fatalf("External initial arc expected to be %s, but was: %s", expectedInitialArc, initialArc)
	}
	_, ok = storyArcs[initialArc]
	if !ok {
		t.Fatalf("External initial arc (%s) expected to be an available key", initialArc)
	}
}

func Test_extractJSONContents(t *testing.T) {
	gopherJSON, err := extractJSONContents("")
	if err != nil {
		t.Fatalf("Failed to extract embeded story arcs: %s", err.Error())
	}

	storyArcs := make(map[string]model.StoryArc)
	err = json.Unmarshal(gopherJSON, &storyArcs)
	if err != nil {
		t.Fatalf("Failed to unmarshal embedded story arcs: %s", err.Error())
	}

	expectedArcs := []string{
		"intro", "new-york", "debate", "sean-kelly",
		"mark-bates", "denver", "home",
	}
	for _, arcName := range expectedArcs {
		_, ok := storyArcs[arcName]
		if !ok {
			t.Fatalf("%s was not found in the embeded story", arcName)
		}
	}

	externalGopherJSON, err := extractJSONContents("gopher-no-intro.json")
	if err != nil {
		t.Fatalf("Failed to extract story arcs from external file: %s", err.Error())
	}

	storyArcs = make(map[string]model.StoryArc)
	err = json.Unmarshal(externalGopherJSON, &storyArcs)
	if err != nil {
		t.Fatalf("Failed to unmarshal external story arcs: %s", err.Error())
	}

	expectedArcs = []string{
		"main", "new-york", "debate", "sean-kelly",
		"mark-bates", "denver", "home",
	}
	for _, arcName := range expectedArcs {
		_, ok := storyArcs[arcName]
		if !ok {
			t.Fatalf("%s was not found in the external story", arcName)
		}
	}
}

func Test_extractInitialArc(t *testing.T) {
	withIntroMap := make(map[string]model.StoryArc)
	expectedInitialArc := "intro"
	withIntroMap[expectedInitialArc] = model.StoryArc{}
	initialArc, err := extractInitialArc(&withIntroMap)
	if err != nil {
		t.Fatalf("Failed to extract initial arc with intro arc: %s", err.Error())
	}
	if initialArc != expectedInitialArc {
		t.Fatalf("Inital arc should have been %s", expectedInitialArc)
	}

	withoutIntroMap := make(map[string]model.StoryArc)
	expectedInitialArc = "main"
	withoutIntroMap[expectedInitialArc] = model.StoryArc{IsInitialStory: true}
	initialArc, err = extractInitialArc(&withoutIntroMap)
	if err != nil {
		t.Fatalf("Failed to extract initial arc without intro arc: %s", err.Error())
	}
	if initialArc != expectedInitialArc {
		t.Fatalf("Inital arc should have been %s", expectedInitialArc)
	}
}
