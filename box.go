package main

import (
	// "math"
	// "fmt"
	"log"
)

type Box struct {
	PrimitiveType
	Coords []int `json:"coords"`

	// Type for now supports "default" or "shadow"
	Type string `json:"type"`
}

func (b *Box) Click(x, y int) {
	if x >= b.Coords[0] &&
		x <= b.Coords[2] &&
		y >= b.Coords[1] &&
		y <= b.Coords[3] {
		b.selected = true
	} else {
		b.selected = false
	}
}

func (b Box) Validate() error {
	//TODO implement
	// - start is lower than end and in bounds
	// - at least 2x2
	return nil
}

func abs(i int) int {
	if i < 0 {
		i = -i
	}
	return i
}

func (b Box) Draw() RuneMap {

	var x1, y1, x2, y2 int = b.Coords[0], b.Coords[1], b.Coords[2], b.Coords[3]

	if x1 >= x2 || y1 >= y2 {
		//TODO better error handling?
		log.Fatalf("Invalid Box coordinates (%vx%v)->(%vx%v)!\n", x1, y1, x2, y2)
	}

	r := RuneMap{}
	r.Set(x1, y1, '┌')
	r.Set(x1, y2, '└')
	r.Set(x2, y1, '┐')
	r.Set(x2, y2, '┘')

	for x := x1+1; x < x2; x++ {
		r.Set(x, y1, '─')
		r.Set(x, y2, '─')

		if b.Type == "shadow" {
			r.Set(x+1, y2+1, '░')
			r.Set(x+2, y2+1, '░')
		}
	}

	for y := y1+1; y < y2; y++ {
		r.Set(x1, y, '│')
		r.Set(x2, y, '│')

		if b.Type == "shadow" {
			r.Set(x2+1, y, '░')
			r.Set(x2+1, y+1, '░')
		}
	}

	return r
}
