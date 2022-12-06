package main

import (
	// "log"
)

// RuneMap represents an infinite plane of runes, addressable by x and y.
type RuneMap struct {
	// data[x][y] = 'r'
	data map[int]map[int]rune
}

func (rm RuneMap) Get(x, y int) rune {
	// log.Println("Requesting", x, y, "from", rm.data)
	if column, ok := rm.data[x]; ok {
		if r, ok := column[y]; ok {
			return r
		}
	}
	return CHAR_EMPTY
}

func (rm *RuneMap) Set(x, y int, r rune) {

	if rm.data == nil {
		rm.data = map[int]map[int]rune{}
	}
	// If the column already exists
	if _, ok := rm.data[x]; ok {
		rm.data[x][y] = r
		return
	}
	// Collumn, did't exist yet, create and set
	rm.data[x] = map[int]rune{y: r}
	rm.data[x][y] = r
}

// MinMax returns the first and last (x, y) coordinate
func (rm RuneMap) MinMax() (int, int, int, int) {

	minX, minY, maxX, maxY := -1, -1, -1, -1

	for x := range rm.data {
		for y := range rm.data[x] {

			if minX == -1 || x < minX {
				minX = x
			}

			if minY == -1 || y < minY {
				minY = y
			}

			if maxX == -1 || x > maxX {
				maxX = x
			}

			if maxY == -1 || y > maxY {
				maxY = y
			}

		}
	}

	if maxX == -1 || maxY == -1 || minX == -1 || minY == -1 {
		return 0, 0, 0, 0
	}
	return minX, minY, maxX, maxY
}

// String represents a RuneMap as string (with newlines)
func (rm RuneMap) String() string {


	out := ""

	minX, minY, maxX, maxY := rm.MinMax()

	if minX == maxX && minY == maxY {
		return ""
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if rm.Get(x, y) == CHAR_EMPTY {
				out = out + " "
				continue
			}
			out = out + string(rm.Get(x, y))
		}
		if y != maxY {
			out = out + "\n"
		}
	}

	return out
}
