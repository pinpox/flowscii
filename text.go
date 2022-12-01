package main

import (
	"strings"
)

type Text struct {
	Coords []int  `json:"coords"`
	Text   string `json:"text"`
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

func (t Text) Drawable() Drawable {

	text_lines := strings.Split(t.Text, "\n")
	//find longest line

	var x int = 1
	for _, v := range text_lines {
		if len(v) > x {
			x = len(v)
		}
	}

	runemap := initRuneMap(x, len(text_lines))
	for ry, line := range text_lines {
		for rx, r := range []rune(line) {
			runemap.Set(rx, ry, r)
		}
	}

	return Drawable{t.Coords[0], t.Coords[1], runemap}
}
