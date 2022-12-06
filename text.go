package main

import (
	"strings"
)

type Text struct {
	PrimitiveType
	Coords []int    `json:"coords"`
	Text   string   `json:"text"`
	Style  []string `json:"style"`
}

func (t Text) isItalic() bool {
	for _, v := range t.Style {
		if v == "italic" {
			return true
		}
	}
	return false
}

func (t Text) isBold() bool {
	for _, v := range t.Style {
		if v == "bold" {
			return true
		}
	}
	return false
}

func (t Text) Draw() RuneMap {

	text_lines := strings.Split(t.Text, "\n")
	//find longest line

	var x int = 1
	for _, v := range text_lines {
		if len(v) > x {
			x = len(v)
		}
	}

	runemap := RuneMap{}
	for ry, line := range text_lines {
		for rx, r := range []rune(line) {
			runemap.Set(t.Coords[0]+rx, t.Coords[1]+ry, r)
		}
	}

	return runemap
}
