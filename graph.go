package main

import (
	"encoding/json"
	"github.com/gdamore/tcell/v2"
	"io/ioutil"
	"log"
	"os"
)

func (g *Graph) AddBox(x1, y1, x2, y2 int) {
	g.Objects.Box = append(g.Objects.Box, Box{
		PrimitiveType{false},
		[]int{x1, y1, x2, y2},
		"default",
	})
}

func (g *Graph) AddLine(coords []int) {
	g.Objects.Line = append(g.Objects.Line, Line{
		PrimitiveType{false},
		coords,
		"default",
	})
}

func (g *Graph) AddText(x1, y1 int) {
	g.Objects.Text = append(g.Objects.Text, Text{
		PrimitiveType{false},
		[]int{x1, y1},
		"NEW TEXT",
		[]string{},
	})
}

func (g *Graph) SaveJSON(path string) {
	//TODO
}

type Graph struct {
	Metadata   Metadata `json:"metadata"`
	Objects    Objects  `json:"objects"`
	events     chan tcell.Event
	ox, oy     int
	oldx, oldy int
}

type Metadata struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (g *Graph) Update() {
	select {
	case event := <-g.events:
		// log.Println("received message in graph", event)
		g.handleEvent(event)
		//TODO Process event
	default:
		log.Println("no message received in graph")
	}
}

func loadGraph(path string) Graph {

	// Open our jsonFile
	jsonFile, err := os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Successfully Opened", path)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var graph Graph

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &graph)

	graph.events = make(chan tcell.Event, 1)

	graph.ox = -1
	graph.oy = -1

	log.Println("Loaded Graph:", graph)

	return graph
}

func (g *Graph) DeselectAll() {
	for k := range g.Objects.Box {
		g.Objects.Box[k].selected = false
	}

	for k := range g.Objects.Line {
		g.Objects.Line[k].selected = false
	}

	for k := range g.Objects.Text {
		g.Objects.Text[k].selected = false
	}
}

type vec2 struct {
	X int
	Y int
}

func (g *Graph) MoveSelected(x, y int) {

	log.Println("MOVE", x, y)
	log.Println("O", g.ox, g.oy)
	log.Println("OLD", g.oldx, g.oldy)

	delta_x := x - g.ox
	delta_y := y - g.oy

	for k, line := range g.Objects.Line {
		if line.Selected() {

			log.Println("Moving Line by:", (delta_x - g.oldx), (delta_y - g.oldy))

			// Check if a corner was clicked
			for i := 0; i < len(line.Coords); i += 2 {
				line_x := g.Objects.Line[k].Coords[i]
				line_y := g.Objects.Line[k].Coords[i+1]

				log.Println("LINE COORDS", line_x, line_y)
				log.Println(g.ox, g.oy, g.oldx, g.oldy, x, y )

				//TODO fix this condition
				if line_x == g.ox + g.oldx && line_y == g.oy +g.oldy {

					log.Println("moving line....", line_x, line_y, delta_x, delta_y)
					log.Println("Moving by:", (delta_x - g.oldx), (delta_y - g.oldy))



					// not first
					if i != 0 {

						if g.Objects.Line[k].Coords[i] == g.Objects.Line[k].Coords[i-2] {
							g.Objects.Line[k].Coords[i-2] += (delta_x - g.oldx)
						}

						if g.Objects.Line[k].Coords[i+1] == g.Objects.Line[k].Coords[i-1] {
							g.Objects.Line[k].Coords[i-1] += (delta_y - g.oldy)
						}
					}

					// not last
					if i != len(line.Coords)-2 {

						if g.Objects.Line[k].Coords[i+2] == g.Objects.Line[k].Coords[i] {
							g.Objects.Line[k].Coords[i+2] += (delta_x - g.oldx)
						}

						if g.Objects.Line[k].Coords[i+1] == g.Objects.Line[k].Coords[i+3] {
							g.Objects.Line[k].Coords[i+3] += (delta_y - g.oldy)
						}
					}

					// 					// Move the corner
					g.Objects.Line[k].Coords[i] += (delta_x - g.oldx)
					g.Objects.Line[k].Coords[i+1] += (delta_y - g.oldy)

					// Cleanup

					// 					// Adjust adjecent ones

					break
					// return

				}

			}
		}
	}

	// Boxes
	for k := range g.Objects.Box {
		if g.Objects.Box[k].Selected() {
			log.Println("Moving Box by:", (delta_x - g.oldx), (delta_y - g.oldy))
			g.Objects.Box[k].Coords[0] += (delta_x - g.oldx)
			g.Objects.Box[k].Coords[2] += (delta_x - g.oldx)
			g.Objects.Box[k].Coords[1] += (delta_y - g.oldy)
			g.Objects.Box[k].Coords[3] += (delta_y - g.oldy)
		}
	}

	// Text
	for k := range g.Objects.Text {
		if g.Objects.Text[k].Selected() {
			log.Println("Moving Text by:", (delta_x - g.oldx), (delta_y - g.oldy))
			g.Objects.Text[k].Coords[0] += (delta_x - g.oldx)
			g.Objects.Text[k].Coords[1] += (delta_y - g.oldy)
		}
	}

	g.oldx = delta_x
	g.oldy = delta_y

}

func (g *Graph) Select(x, y int) {

	// Text
	// If text is clicked, select only that
	for k, v := range g.Objects.Text {
		if v.Draw().Get(x, y) != CHAR_EMPTY {
			g.Objects.Text[k].selected = true
			return
		}
	}
	// Line
	for k, v := range g.Objects.Line {
		if v.Draw().Get(x, y) != CHAR_EMPTY {
			g.Objects.Line[k].selected = true
			return
		}
	}

	// Box
	// Select indices of all boxes clicked inside
	var boxes_sel []int
	for k, v := range g.Objects.Box {
		if x >= v.Coords[0] && x <= v.Coords[2] && y >= v.Coords[1] && y <= v.Coords[3] {
			boxes_sel = append(boxes_sel, k)
		}
	}

	// Find smallest most inner (if any)
	if len(boxes_sel) >= 1 {

		// select first box
		var box_n int
		box_n = boxes_sel[0]

		// (x1_old, y1_old)
		//        +──────────────────────────────┐
		//        │                              │
		//        │(x1_new, y1_new)              │
		//        │       +──────────┐           │
		//        │       │          │           │
		//        │       │          │           │
		//        │       └──────────+           │
		//        │           (x2_new, y2_new)   │
		//        │                              │
		//        └──────────────────────────────+
		//                                (x2_old, y2_old)

		// x - x1_new <= x - x1_old || x2_new - x < x2_old - x ||
		// y - y1_new <= y - y1_old || y2_new - y < y2_old - y

		// Find smallest box around x, y
		for _, v := range boxes_sel {

			coords_old := g.Objects.Box[box_n].Coords
			coords_new := g.Objects.Box[v].Coords

			if x-coords_new[0] <= x-coords_old[0] || coords_new[2]-x < coords_old[2]-x ||
				y-coords_new[1] <= y-coords_old[1] || coords_new[3]-y < coords_old[3]-y {
				box_n = v
			}
		}

		// Also find all boxes that are inside it
		boxes_sel = []int{box_n}
		c := g.Objects.Box[box_n].Coords

		for k, v := range g.Objects.Box {
			if (v.Coords[0] >= c[0] && v.Coords[0] <= c[2]) && (v.Coords[1] >= c[1] && v.Coords[1] <= c[3]) &&
				(v.Coords[2] >= c[0] && v.Coords[2] <= c[2]) && (v.Coords[3] >= c[1] && v.Coords[3] <= c[3]) {
				boxes_sel = append(boxes_sel, k)
			}
		}

		// Also find text that is inside that box
		texts_sel := []int{}

		for k, v := range g.Objects.Text {
			if (v.Coords[0] >= c[0] && v.Coords[0] <= c[2]) && (v.Coords[1] >= c[1] && v.Coords[1] <= c[3]) {
				texts_sel = append(texts_sel, k)
			}
		}
		// TODO find lines inside the box

		for _, v := range boxes_sel {
			g.Objects.Box[v].selected = true
		}

		for _, v := range texts_sel {
			g.Objects.Text[v].selected = true
		}

	}

}

func (g *Graph) handleEvent(ev tcell.Event) {

	switch ev := ev.(type) {
	case *tcell.EventKey:
		// if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
		//	return
		// } else if ev.Key() == tcell.KeyCtrlL {
		//	s.Sync()
		// } else if ev.Rune() == 'C' || ev.Rune() == 'c' {
		//	s.Clear()
		// }
	case *tcell.EventMouse:
		x, y := ev.Position()

		// log.Println(x)

		switch ev.Buttons() {
		case tcell.Button1, tcell.Button2:
			if g.ox < 0 {
				g.ox, g.oy = x, y // record location when click started
			}

			if g.ox == x && g.oy == y {
				g.Select(x, y)
			}
			g.MoveSelected(x, y)

		case tcell.ButtonNone:
			if g.ox >= 0 {

				g.DeselectAll()
				g.oldx = 0
				g.oldy = 0

				// msg := "hi"
				// log.Printf("GRAPH Dragged: %d,%d to %d,%d", g.ox, g.oy, x, y)
				g.ox, g.oy = -1, -1
			}
		}
	}

}

type Objects struct {
	Box  []Box  `json:"box"`
	Line []Line `json:"line"`
	Text []Text `json:"text"`
}

type PrimitiveType struct {
	selected bool
}

func (p PrimitiveType) Selected() bool {
	return p.selected
}

type Primitive interface {
	Draw() RuneMap
	Selected() bool
	// Click(x, y int)
	// Drag(x1, y1, x2, y2 int)
	// TODO
	// Validate() error
}
