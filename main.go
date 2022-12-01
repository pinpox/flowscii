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
	Metadata Metadata `json:"metadata"`
	Objects  Objects  `json:"objects"`
}

type Metadata struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type Objects struct {
	Box  []Box  `json:"box"`
	Line []Line `json:"line"`
	Text []Text `json:"text"`
}

type Text struct {
	Coords []int  `json:"coords"`
	Text   string `json:"text"`
}

type Primitive interface {
	Drawable() Drawable
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

	log.Println("Loaded Graph:", graph)

	return graph
}

func main2() {
	var graph Graph = loadGraph("example.json")

	var canvas Canvas

	for _, v := range graph.Objects.Box {
		// d := v.Drawable()
		// fmt.Printf("Box at (%v,%v):\n%v\n", d.StartX, d.StartY, d.Content.String())
		canvas.Add(v)
	}

	for _, v := range graph.Objects.Line {
		// d := v.Drawable()
		// fmt.Printf("Line at (%v,%v):\n%v\n", d.StartX, d.StartY, d.Content.String())
		canvas.Add(v)
	}

	// for _, v := range graph.Objects.Line {
	//	// TODO draw lines
	//	x, y, text := v.Drawble()
	//	canvas.Add(v)
	// }

	// for _, v := range graph.Objects.Text {
	//	// TODO draw text
	//	x, y, text := v.Drawable()
	//	canvas.Add(v)
	// }

	// fmt.Println("Resulting Canvas:")
	fmt.Println(canvas.String())

}

// START TCELL

func drawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func drawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// Fill background
	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			s.SetContent(col, row, ' ', nil, style)
		}
	}

	// Draw borders
	for col := x1; col <= x2; col++ {
		s.SetContent(col, y1, tcell.RuneHLine, nil, style)
		s.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}

	// Only draw corners if necessary
	if y1 != y2 && x1 != x2 {
		s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
		s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
		s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
		s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
	}

	drawText(s, x1+1, y1+1, x2-1, y2-1, style, text)
}

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

func drawGPrimitive(s tcell.Screen, v Primitive) {
	log.Println("drawing primitive:", v)
	d := v.Drawable()
	dimX, dimY := d.Content.Dims()

	style := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

	for x := 0; x < dimX; x++ {
		for y := 0; y < dimY; y++ {
			log.Println("Drawing at", x, y)

			current, _, _, _ := s.GetContent(x+d.StartX, y+d.StartY)
			// TODO replacement rule for line joins
			if d.Content.Get(x, y) == '.' || d.Content.Get(x, y) == current {
				continue
			}

			if current == tcell.RuneHLine || current == tcell.RuneVLine || current == tcell.RuneTTee || current == tcell.RuneRTee || current == tcell.RuneLTee || current == tcell.RuneBTee || current == tcell.RuneULCorner || current == tcell.RuneURCorner || current == tcell.RuneURCorner || current == tcell.RuneLLCorner || current == tcell.RuneLRCorner {
				s.SetContent(x+d.StartX, y+d.StartY, tcell.RunePlus, nil, style)
				continue
			}

			s.SetContent(x+d.StartX, y+d.StartY, d.Content.Get(x, y), nil, style)
		}
	}

}

func drawGraph(s tcell.Screen, g Graph) {

	for _, v := range g.Objects.Box {
		drawGPrimitive(s, v)
	}

	for _, v := range g.Objects.Line {
		drawGPrimitive(s, v)
	}

	// TODO implement text primitive
	// for _, v := range g.Objects.Text {
	// drawGPrimitive(s, v)
	// }

	// Clean up junctions
	xmax, ymax := s.Size()

	style := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

	for x := 0; x < xmax; x++ {
		for y := 0; y < ymax; y++ {

			c, _, _, _ := s.GetContent(x, y)

			if c == tcell.RunePlus {

				// No vertical line above
				c_above, _, _, _ := s.GetContent(x, y-1)
				if !(c_above == tcell.RuneVLine) {
					s.SetContent(x, y, tcell.RuneTTee, nil, style)
				}

				// No vertical line below
				c_below, _, _, _ := s.GetContent(x, y+1)
				if !(c_below == tcell.RuneVLine) {
					s.SetContent(x, y, tcell.RuneBTee, nil, style)
				}

				// No horizontal line right
				c_right, _, _, _ := s.GetContent(x+1, y)
				if !(c_right == tcell.RuneHLine) {
					s.SetContent(x, y, tcell.RuneRTee, nil, style)
				}

				// No horizontal line left
				c_left, _, _, _ := s.GetContent(x-1, y)
				if !(c_left == tcell.RuneHLine){
					s.SetContent(x, y, tcell.RuneLTee, nil, style)
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
	boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorGrey)

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
		s.Show()

		// Poll event
		ev := s.PollEvent()

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

			switch ev.Buttons() {
			case tcell.Button1, tcell.Button2:
				if ox < 0 {
					ox, oy = x, y // record location when click started
				}

			case tcell.ButtonNone:
				if ox >= 0 {
					label := fmt.Sprintf("%d,%d to %d,%d", ox, oy, x, y)
					drawBox(s, ox, oy, x, y, boxStyle, label)
					ox, oy = -1, -1
				}
			}
		}
	}
}
