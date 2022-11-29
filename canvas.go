package main

import (
	"fmt"
	// "reflect"
)

type Drawable struct {
	StartX  int
	StartY  int
	Content RuneMap
}

type Canvas struct {
	Primitives []Primitive
	// startX, startY
	Coords []int
}

func (c Canvas) String() string {
	//TODO

	var drawables []Drawable

	// Render drawables
	for _, v := range c.Primitives {
		drawables = append(drawables, v.Drawable())
	}

	// Find lowest coordinate
	var minX, minY int
	for _, v := range drawables {
		if v.StartX < minX {
			minX = v.StartX
		}
		if v.StartY < minY {
			minY = v.StartY
		}
	}

	// Find highest coordinate
	var maxX, maxY int
	for _, v := range drawables {
		if v.StartX+len(v.Content) > maxX {
			maxX = v.StartX + len(v.Content)
		}
		if v.StartY+len(v.Content[0]) > maxY {
			maxY = v.StartY+len(v.Content[0])
		}
	}

	fmt.Println("Min X/Y", minX, minY)
	fmt.Println("Max X/Y", maxX, maxY)

	var r RuneMap = make([][]rune, maxX)
	for i := 0; i < maxX; i++ {
		r[i] = make([]rune, maxY)
	}






	for x := range r {
		for y := range r[x] {
			r[x][y] = '.'
		}
	}

	for _, d := range drawables {
		for x := range d.Content {
			for y := range d.Content[x] {
				// TODO add overlapping/replacing rules
				r[x+d.StartX][y+d.StartY] = d.Content[x][y]
			}
		}
	}


	return r.String()

}

func (c *Canvas) Add(p Primitive) {
	c.Primitives = append(c.Primitives, p)
}
