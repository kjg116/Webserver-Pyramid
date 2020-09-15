package models

// Request ..
type Request struct {
	Words []string `json:"words"`
}

// Result ..
type Result struct {
	CharacterMapSlice CharacterMapSlice `json:"characters,omitempty"`
	IsPyramid         bool              `json:"is_pyramid"`
	Word              string            `json:"word"`
	ShowDetails       bool
}

// Response ..
type Response struct {
	Results []Result `json:"results"`
}
