package main

type Box struct {
	Coords []int `json:"coords"`

	// Boxtype for now supports "default" or "shadow"
	Boxtype string `json:"boxtype"`
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


func (b Box) Drawable() Drawable {

	var out RuneMap

	length := b.Coords[2] - b.Coords[0]
	height := b.Coords[3] - b.Coords[1]

	out = append(out, makeRow('┌', '─', '┐', length))
	for i := 0; i < height-2; i++ {
		out = append(out, makeRow('│', ' ', '│', length))
	}
	out = append(out, makeRow('└', '─', '┘', length))

	if b.Boxtype == "shadow" {
		out[0] = append(out[0], ' ')
		for i := 1; i < height; i++ {
			out[i] = append(out[i], '░')
		}
		out = append(out, makeRow(' ', '░', '░', length+1))
	}

	return Drawable{b.Coords[0], b.Coords[1], out}
}
