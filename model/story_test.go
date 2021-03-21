package model_test

import (
	"embed"
	"testing"

	"github.com/ProFL/gophercises-cyoa/model"
)

//go:embed fixtures/gopher.json
var fixturesFs embed.FS

func Test_ParseStory_EmbeddedStory(t *testing.T) {
	expectedInitialArc := "intro"
	story := model.ParseStory("fixtures/gopher.json", &fixturesFs)
	if story.InitialArc != expectedInitialArc {
		t.Fatalf("Embedded initial arc expected to be %s, but was: %s",
			expectedInitialArc, story.InitialArc)
	}
	_, ok := story.Arcs[story.InitialArc]
	if !ok {
		t.Fatalf("Embedded initial arc (%s) expected to be an available key", story.InitialArc)
	}
}

func Test_ParseStory_ExternalStory(t *testing.T) {
	expectedInitialArc := "initial"
	story := model.ParseStory("fixtures/story-no-intro.json", nil)
	if story.InitialArc != expectedInitialArc {
		t.Fatalf("External initial arc expected to be %s, but was: %s",
			expectedInitialArc, story.InitialArc)
	}
	_, ok := story.Arcs[story.InitialArc]
	if !ok {
		t.Fatalf("External initial arc (%s) expected to be an available key", story.InitialArc)
	}
}
