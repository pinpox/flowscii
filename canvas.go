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

		dX, dY := v.Content.Dims()

		if v.StartX+dX > maxX {
			maxX = v.StartX + dX
		}
		if v.StartY+dY > maxY {
			maxY = v.StartY + dY
		}
	}

	fmt.Println("Min X/Y", minX, minY)
	fmt.Println("Max X/Y", maxX, maxY)

	var r RuneMap = initRuneMap(maxX, maxY)

	for _, d := range drawables {

		dX, dY := d.Content.Dims()

		for x := 0; x < dX; x++ {
			for y := 0; y < dY; y++ {
				if d.Content.Get(x, y) == '.' {
					continue
				}
				r.Set(x+d.StartX, y+d.StartY, d.Content.Get(x, y))
			}
		}
	}

	return fmt.Sprint(r)

}

func (c *Canvas) Add(p Primitive) {
	c.Primitives = append(c.Primitives, p)
}
