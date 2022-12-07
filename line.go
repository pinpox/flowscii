package main

import (
	"log"
)


type Line struct {
	PrimitiveType
	Coords []int `json:"coords"`

	// Type supports "default" and "arrow"
	Type string `json:"type"`
}

func normalizeCoords(c []int) (coords []int, offsetX int, offsetY int) {
	coords = append(coords, c...)
	//find minimum X and Y
	minX, minY := c[0], c[1]

	for i := 0; i < len(c); i += 2 {
		if c[i] < minX {
			minX = c[i]
		}
		if c[i+1] < minY {
			minY = c[i+1]
		}
	}

	for i := 0; i < len(c); i += 2 {
		coords[i] = c[i] - minX
		coords[i+1] = c[i+1] - minY
	}
	return coords, minX, minY
}

func (l Line) Draw() RuneMap{


		log.Printf("Line Coords: %v", l.Coords)

	r := RuneMap{}
	coords := l.Coords

	// "coords": [ 11,11, 11,15, 5,15 ],

	//draw lines
	// fmt.Println(r)

	var prevDir int = 0

	for i := 2; i < len(coords); i += 2 {
		var prevX, prevY int = coords[i-2], coords[i-1]

		//    1
		//  4 0 2
		//    3

		var x, y int = coords[i], coords[i+1]

		// dimRX, dimRY := r.Dims()
		// fmt.Printf("drawing line (%v,%v) -> (%v,%v) on map with dims: %vx%v\n", prevX, prevY, x, y, dimRX, dimRY)

		// Vertical line
		if x == prevX {

			// Down
			if prevY < y {
				for i := prevY; i <= y; i++ {
					r.Set(x, i, '│')
				}

				switch prevDir {
				case 2:
					r.Set(x, prevY, '┐')
				case 4:
					r.Set(x, prevY, '┌')
				}

				prevDir = 3
			} else {
				// Up
				for i := y; i <= prevY; i++ {
					r.Set(x, i, '│')
				}

				switch prevDir {
				case 2:
					r.Set(x, prevY, '┘')
				case 4:
					r.Set(x, prevY, '└')
				}

				prevDir = 1
			}

			continue
		}

		// Horizontal line
		if y == prevY {

			if prevX < x {
				// fmt.Println("drawing right")
				// right
				for i := prevX; i <= x; i++ {
					r.Set(i, y, '─')
				}

				switch prevDir {
				case 1:
					r.Set(prevX, y, '┌')
				case 3:
					r.Set(prevX, y, '└')
				}

				prevDir = 2
			} else {
				// fmt.Println("drawing left")
				// left
				for i := x; i <= prevX; i++ {
					r.Set(i, y, '─')
				}
				switch prevDir {
				case 1:
					r.Set(prevX, y, '┐')
				case 3:
					r.Set(prevX, y, '┘')
				}
				prevDir = 4
			}

			continue
		}
		log.Fatalf("Invalid Line %v", coords)

	}

	arrows := []rune{'▲', '►', '▼', '◄'}
	if l.Type == "double_arrow" || l.Type == "arrow" {
		r.Set(coords[len(coords)-2], coords[len(coords)-1], arrows[prevDir-1])
	}

	if l.Type == "double_arrow" {

		if coords[0] > coords[2] {
			r.Set(coords[0], coords[1], arrows[1])
		}

		if coords[0] < coords[2] {
			r.Set(coords[0], coords[1], arrows[3])
		}

		if coords[1] > coords[3] {
			r.Set(coords[0], coords[1], arrows[2])
		}

		if coords[1] < coords[3] {
			r.Set(coords[0], coords[1], arrows[1])
		}

	}

	return r

}
