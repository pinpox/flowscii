package main

import (
	// "math"
	// "fmt"
	"log"
)

type Box struct {
	Coords []int `json:"coords"`

	// Type for now supports "default" or "shadow"
	Type string `json:"type"`
}

func makeRow(start, mid, end rune, length int) []rune {
	var out []rune = []rune{start}

	for i := 0; i < length-2; i++ {
		out = append(out, mid)
	}
	out = append(out, end)
	return out
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

func (b Box) Drawable() Drawable {

	var x1, y1, x2, y2 int = b.Coords[0], b.Coords[1], b.Coords[2], b.Coords[3]

	if x1 >= x2 || y1 >= y2 {
		//TODO better error handling?
		log.Fatalf("Invalid Box coordinates (%vx%v)->(%vx%v)!\n", x1, y1, x2, y2)
	}

	lenX := x2 - x1 + 1
	lenY := y2 - y1 + 1

	offsetX := x1
	offsetY := y1

	x2 = x2-x1
	y2 = y2-y1
	x1, y1 = 0, 0

	r := initRuneMap(lenX, lenY)
	r.Set(lenX-1, 0, '┐')
	r.Set(lenX-1, lenY-1, '┘')
	r.Set(0, lenY-1, '└')
	r.Set(0, 0, '┌')

	for x := x1 + 1; x < x2; x++ {
		r.Set(x, y1, '─')
		r.Set(x, y2, '─')
	}

	for y := y1 + 1; y < y2; y++ {
		r.Set(x1, y, '│')
		r.Set(x2, y, '│')
	}


	if b.Type == "shadow" {

		rnew := initRuneMap(lenX+1, lenY+1)

		for x := 0; x <lenX; x++ {
			for y := 0; y <lenY; y++ {
				rnew.Set(x,y, r.Get(x,y))
			}
		}

		// Vertical shadow
		for i := 1; i < lenY; i++ {
			rnew.Set( lenX, i, '░')
		}

		rnew.data[lenY] =  makeRow('.', '░', '░',lenX+1)

		r = rnew
	}

	return Drawable{offsetX, offsetY, r}
}
