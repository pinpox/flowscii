package main

import (
	// "math"
	// "fmt"
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

	// Coords: []int{2, 0, 0, 2},

	// normalize: (x1, y1) bottomleft, (x2, y2) top right

	var x1, x2, y1, y2 int

	if b.Coords[2] > b.Coords[0] {
		x1 = b.Coords[0]
		x2 = b.Coords[2]
	} else {
		x1 = b.Coords[2]
		x2 = b.Coords[0]
	}

	if b.Coords[3] > b.Coords[1] {
		y1 = b.Coords[1]
		y2 = b.Coords[3]
	} else {
		y1 = b.Coords[3]
		y2 = b.Coords[1]
	}

	// fmt.Printf("Drawing Box (%v,%v) -> (%v,%v)", x1, y1, x2, y2)

	lenX := x2 - x1 + 1
	lenY := y2 - y1 + 1

	offsetX := x1
	offsetY := y1


	x2 = x2-x1
	y2 = y2-y1
	x1, y1 = 0, 0

	r := initRuneMap(lenX, lenY)

	// fmt.Println("len x/y", lenX, lenY)
	// dx, dy := r.Dims()
	// fmt.Println("dims", dx, dy)

	r.Set(0, 0, '└')
	r.Set(lenX-1, 0, '┘')
	r.Set(lenX-1, lenY-1, '┐')
	r.Set(0, lenY-1, '┌')

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
				rnew.Set(x,y+1, r.Get(x,y))
			}
		}

		// rnew.Set(0, lenY, '.')

		// Vertical shadow
		for i := 1; i < lenY; i++ {
			rnew.Set( lenX, i, '░')
		}


		rnew.data[0] =  makeRow('.', '░', '░',lenX+1)

		r = rnew
	}

	return Drawable{offsetX, offsetY, r}
}
