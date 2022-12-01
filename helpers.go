package main
// import ("log")


type RuneMap struct {
	data [][]rune
}

func (rm RuneMap) Get(x, y int) rune {
	// log.Println("Requesting", x, y, "from", rm.data)
	return rm.data[y][x]
}

func (rm RuneMap) Set(x, y int, r rune) {
	rm.data[y][x] = r
}

func (rm RuneMap) Dims() (int, int) {
	return len(rm.data[0]), len(rm.data)
}

// Represent [][]rune as string (with newlines)
func (rm RuneMap) String() string {

	// start at 0, maxY
	// go to maxX, maxY
	// go to 0, maxY-1

	out := ""

	lenX, lenY := rm.Dims()

	for y := lenY - 1; y >= 0; y-- {

		for x := 0; x <lenX; x++ {
			out = out + string(rm.Get(x, y))
		}

		if y != 0 {
			out = out + "\n"
		}

	}

	return out

}

func initRuneMap(x, y int) RuneMap {
	var r [][]rune = make([][]rune, y)
	for i := 0; i < y; i++ {
		r[i] = make([]rune, x)
	}

	for y := range r {
		for x := range r[y] {
			r[y][x] = '.'
		}
	}

	// fmt.Printf("X: %v, Y: %v Created %v" , x, y, r)
	return RuneMap{r}
}
