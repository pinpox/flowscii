package main

import (
	"log"
	"os"
	"strings"
	"github.com/gdamore/tcell/v2"
)

const CHAR_EMPTY rune = '\x00'

// Map function over slice
func map2[T, U any](data []T, f func(T) U) []U {
	res := make([]U, 0, len(data))
	for _, e := range data {
		res = append(res, f(e))
	}
	return res
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
			if d.Content.Get(x, y) == CHAR_EMPTY || d.Content.Get(x, y) == current {
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
