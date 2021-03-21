package model

type StoryArc struct {
	Title      string           `json:"title"`
	Paragraphs []string         `json:"story"`
	Options    []StoryArcOption `json:"options"`
}
