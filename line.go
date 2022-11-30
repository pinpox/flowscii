package main

import "fmt"

type Line struct {
	Coords []int `json:"coords"`

	// Type supports "default" and "arrow"
	Type string `json:"type"`
}

func findMinMaxCoords(c []int) (minX, minY, maxX, maxY int) {

	minX, maxX = c[0], c[0]
	minY, maxY = c[1], c[1]

	for i := 0; i < len(c); i += 2 {

		if c[i] > maxX {
			maxX = c[i]
		}

		if c[i] < minX {
			minX = c[i]
		}

		if c[i+1] > maxY {
			maxY = c[i+1]
		}

		if c[i+1] < minY {
			minY = c[i+1]
		}

	}

	return minX, minY, maxX, maxY

}

func (l Line) Drawable() Drawable {

	//find min/max and create box
	minX, minY, maxX, maxY := findMinMaxCoords(l.Coords)
	r := initRuneMap(maxX+1-minX, maxY+1-minY)

	//draw lines
	// fmt.Println(r)

	var prevDir int = 0
	for i := 2; i < len(l.Coords); i += 2 {
		var prevX, prevY int = l.Coords[i-2], l.Coords[i-1]

		//    1
		//  4 0 2
		//    3

		var x, y int = l.Coords[i], l.Coords[i+1]

		dimRX, dimRY := r.Dims()
		fmt.Printf("drawing line (%v,%v) -> (%v,%v) on map with dims: %vx%v\n", prevX, prevY, x, y, dimRX, dimRY)

		// Vertical line
		if x == prevX {

			// ┌─────────────┐                   ┌─────────────┐
			// │   ROUTER    │░ 192.168.45.0/24  │   HA-PORT   │░
			// INTERNET──►│ inovex-prod ├──────────────────►│    Dummy    │░
			// │             │░                  └──────┬──────┘░
			// └─────────────┘░                   ░░░░░░│░░░░░░░░
			//  ░░░░░░░░░░░░░░░

			// draw line from (x, prevY) to (x, y)
			if prevY < y {
				fmt.Println("drawing up")
				// Up
				for i := prevY; i <= y; i++ {
					r.Set(x, i, '│')
				}

				fmt.Println("Prev was ", prevDir)
				switch prevDir {
				case 2:
					r.Set(x, prevY, '┘')
				case 4:
					r.Set(x, prevY, '└')
				}

				prevDir = 1
			} else {
				fmt.Println("drawing down")
				//Down
				for i := y; i <= prevY; i++ {
					r.Set(x, i, '│')
				}

				switch prevDir {
				case 2:
					r.Set(x, prevY, '┐')
				case 4:
					r.Set(x, prevY, '┌')
				}

				prevDir = 3
			}
			continue
		}

		// Horizontal line
		if y == prevY {

			if prevX < x {
				fmt.Println("drawing right")
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
				fmt.Println("drawing left")
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
		panic("Invalid line coords")

	}

	return Drawable{minX, minY, r}
}
