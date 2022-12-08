package main

import (
	"log"
	// "encoding/json"
)

type Line struct {
	PrimitiveType
	Coords []Vec2 `json:"coords"`
	coord_selected int

	// Type supports "default" and "arrow"
	Type string `json:"type"`
}

func (l Line) Draw() RuneMap {

	// log.Printf("Line Coords: %v", l.Coords)

	r := RuneMap{}

	var prevDir int = 0

	for i := 1; i < len(l.Coords); i++ {

		var coord = l.Coords[i]
		var coord_prev = l.Coords[i-1]

		//    1
		//  4 0 2
		//    3

		var x, y int = coord.X, coord.Y
		var prevX, prevY int = coord_prev.X, coord_prev.Y

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
		log.Fatalf("Invalid Line %v", l.Coords)

	}

	arrows := []rune{'▲', '►', '▼', '◄'}
	if l.Type == "double_arrow" || l.Type == "arrow" {
		r.Set(
			l.Coords[len(l.Coords)-1].X,
			l.Coords[len(l.Coords)-1].Y,
			arrows[prevDir-1])
	}

	if l.Type == "double_arrow" {

		if l.Coords[0].X > l.Coords[1].X {
			r.Set(l.Coords[0].X, l.Coords[0].Y, arrows[1])
		}

		if l.Coords[0].X > l.Coords[1].X {
			r.Set(l.Coords[0].X, l.Coords[0].Y, arrows[3])
		}

		if l.Coords[0].Y > l.Coords[1].Y {
			r.Set(l.Coords[0].X, l.Coords[0].Y, arrows[2])
		}

		if l.Coords[0].Y < l.Coords[1].Y {
			r.Set(l.Coords[0].X, l.Coords[0].Y, arrows[1])
		}

	}

	return r

}
