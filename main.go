package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"log"

	"github.com/gdamore/tcell/v2"
)

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
		log.Println("received message in graph", event)
		g.handleEvent(event)
		//TODO Process event
	default:
		log.Println("no message received in graph")
	}
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

func (g *Graph) MoveSelected(delta_x, delta_y int) {

	for k := range g.Objects.Box {
		if g.Objects.Box[k].Selected() {
			g.Objects.Box[k].Coords[0] += (delta_x - g.oldx)
			g.Objects.Box[k].Coords[2] += (delta_x - g.oldx)
			g.Objects.Box[k].Coords[1] += (delta_y - g.oldy)
			g.Objects.Box[k].Coords[3] += (delta_y - g.oldy)
		}
	}

	g.oldx = delta_x
	g.oldy = delta_y

}

func (g *Graph) Select(x, y int) {

	// If text is clicked, select only that

	for k, v := range g.Objects.Text {

		d := v.Drawable()
		dimX, dimY := d.Content.Dims()

		if x >= v.Drawable().StartX && x <= d.StartX+dimX && y >= d.StartY && y <= d.StartY+dimY {
			g.Objects.Text[k].selected = true
			return
		}
	}

	// if x >= v.Coords[0] && x <= v.Coords[2] && y >= v.Coords[1] && y <= v.Coords[3] {
	// }

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

		for _, v := range boxes_sel {
			g.Objects.Box[v].selected = true
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
			g.MoveSelected(x-g.ox, y-g.oy)

			log.Printf("GRAPH DDDDD: %d,%d to %d,%d", g.ox, g.oy, x, y)

		case tcell.ButtonNone:
			if g.ox >= 0 {

				g.DeselectAll()
				g.oldx = 0
				g.oldy = 0

				// msg := "hi"

				log.Printf("GRAPH Dragged: %d,%d to %d,%d", g.ox, g.oy, x, y)
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
	commands chan PrimitiveCommand
	selected bool
}

func (p PrimitiveType) Selected() bool {
	return p.selected
}

type Primitive interface {
	Drawable() Drawable
	Selected() bool
	// Click(x, y int)
	// Drag(x1, y1, x2, y2 int)
	// TODO
	// Validate() error
}

// Map function over slice
func map2[T, U any](data []T, f func(T) U) []U {
	res := make([]U, 0, len(data))
	for _, e := range data {
		res = append(res, f(e))
	}
	return res
}

func loadGraph(path string) Graph {

	// Open our jsonFile
	jsonFile, err := os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
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

	for k := range graph.Objects.Box {
		graph.Objects.Box[k].commands = make(chan PrimitiveCommand, 1)
	}

	graph.ox = -1
	graph.oy = -1

	log.Println("Loaded Graph:", graph)

	return graph
}

type PrimitiveCommandType int64

const (
	SELECT PrimitiveCommandType = iota
	// DELETE
	// RESIZE
	// UPDATE
)

type PrimitiveCommand struct {
	Type   PrimitiveCommandType
	Params []int
}

// START TCELL

// -----> X
// |
// |
// |
// V  Y

func drawBar(s tcell.Screen, style tcell.Style, items []string) {

	row := 0
	col := 0

	text := []rune(strings.Join(items, " | "))
	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++
	}

	xmax, _ := s.Size()
	for i := col; i < xmax; i++ {
		s.SetContent(i, row, ' ', nil, style)
	}
}

func drawGPrimitive(s tcell.Screen, v Primitive, style tcell.Style) {
	// log.Println("drawing primitive:", v)
	d := v.Drawable()
	dimX, dimY := d.Content.Dims()

	if v.Selected() {
		log.Println("DRAWING SELECTED")
		style = style.Foreground(tcell.ColorMediumVioletRed)
	}

	for x := 0; x < dimX; x++ {
		for y := 0; y < dimY; y++ {
			// log.Println("Drawing at", x, y)

			current, _, _, _ := s.GetContent(x+d.StartX, y+d.StartY)
			// TODO replacement rule for line joins
			if d.Content.Get(x, y) == '.' || d.Content.Get(x, y) == current {
				continue
			}

			if current == tcell.RuneHLine ||
				current == tcell.RuneVLine ||
				current == tcell.RuneTTee ||
				current == tcell.RuneRTee ||
				current == tcell.RuneLTee ||
				current == tcell.RuneBTee ||
				current == tcell.RuneULCorner ||
				current == tcell.RuneURCorner ||
				current == tcell.RuneLLCorner ||
				current == tcell.RuneLRCorner {
				s.SetContent(x+d.StartX, y+d.StartY, tcell.RunePlus, nil, style)
				continue
			}

			s.SetContent(x+d.StartX, y+d.StartY, d.Content.Get(x, y), nil, style)
		}
	}

}

func drawGraph(s tcell.Screen, g Graph) {

	// log.Printf("Drawing graph:\n%+v", g)

	style := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

	for _, v := range g.Objects.Box {
		drawGPrimitive(s, &v, style)
	}

	for _, v := range g.Objects.Line {
		drawGPrimitive(s, v, style)
	}

	for _, v := range g.Objects.Text {
		drawGPrimitive(s, v, style.Bold(v.isBold()).Italic(v.isItalic()))
	}

	// Clean up junctions
	xmax, ymax := s.Size()

	for x := 0; x < xmax; x++ {
		for y := 0; y < ymax; y++ {

			c, _, stylec, _ := s.GetContent(x, y)

			if c == tcell.RunePlus {

				// No vertical line above
				c_above, _, _, _ := s.GetContent(x, y-1)
				if !(c_above == tcell.RuneVLine) {
					s.SetContent(x, y, tcell.RuneTTee, nil, stylec)
				}

				// No vertical line below
				c_below, _, _, _ := s.GetContent(x, y+1)
				if !(c_below == tcell.RuneVLine) {
					s.SetContent(x, y, tcell.RuneBTee, nil, stylec)
				}

				// No horizontal line right
				c_right, _, _, _ := s.GetContent(x+1, y)
				if !(c_right == tcell.RuneHLine) {
					s.SetContent(x, y, tcell.RuneRTee, nil, stylec)
				}

				// No horizontal line left
				c_left, _, _, _ := s.GetContent(x-1, y)
				if !(c_left == tcell.RuneHLine) {
					s.SetContent(x, y, tcell.RuneLTee, nil, stylec)
				}
			}

		}

	}

}

func main() {

	// Init logging
	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	//load graph
	var graph Graph = loadGraph("example.json")

	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	styleBar := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorDarkBlue)

	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.EnableMouse()
	s.EnablePaste()
	s.Clear()

	// Draw initial boxes
	// drawBox(s, 1, 1, 42, 7, boxStyle, "Click and drag to draw a box")
	// drawBox(s, 5, 9, 32, 14, boxStyle, "Press C to reset")

	// drawText(s, 0, 0, 100, 10, styleBar, "TEST")
	drawBar(s, styleBar, []string{"[s] save", "[m] metadata", "[?] help", "[esc] quit"})
	drawGraph(s, graph)

	quit := func() {
		// You have to catch panics in a defer, clean up, and
		// re-raise them - otherwise your application can
		// die without leaving any diagnostic trace.
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	// Here's how to get the screen size when you need it.
	// xmax, ymax := s.Size()

	// Here's an example of how to inject a keystroke where it will
	// be picked up by the next PollEvent call.  Note that the
	// queue is LIFO, it has a limited length, and PostEvent() can
	// return an error.
	// s.PostEvent(tcell.NewEventKey(tcell.KeyRune, rune('a'), 0))

	// Event loop
	ox, oy := -1, -1
	for {

		// Update screen
		s.Clear()

		drawBar(s, styleBar, []string{"[s] save", "[m] metadata", "[?] help", "[esc] quit"})
		drawGraph(s, graph)
		s.Show()

		// Poll event
		ev := s.PollEvent()

		select {
		case graph.events <- ev:
			log.Println("sent message to grarph")
		default:
			log.Println("no message sent to graph")
		}
		graph.Update()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				return
			} else if ev.Key() == tcell.KeyCtrlL {
				s.Sync()
			} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
				s.Clear()
			}
		case *tcell.EventMouse:
			x, y := ev.Position()

			// log.Println(x)

			switch ev.Buttons() {
			case tcell.Button1, tcell.Button2:
				if ox < 0 {
					ox, oy = x, y // record location when click started
				}

				log.Printf("Dragged: %d,%d to %d,%d", ox, oy, x, y)

			case tcell.ButtonNone:
				if ox >= 0 {

					// msg := "hi"

					log.Printf("Dragged: %d,%d to %d,%d", ox, oy, x, y)
					ox, oy = -1, -1
				}
			}
		}
	}
}
