package models

type StoryArc struct {
	Title      string           `json:"title"`
	Paragraphs []string         `json:"story"`
	Options    []StoryArcOption `json:"options"`
}
