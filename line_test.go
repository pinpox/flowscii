package main

import (
	"reflect"
	"testing"
)

func TestLine_Drawable(t *testing.T) {


// -----> X
// |
// |
// |
// V  Y

	rm_tr := RuneMap{[][]rune{{'.', '.', '.'}, {'.', '.', '.'}, {'.', '.', '.'}}}
	rm_tr.Set(0, 0, '─')
	rm_tr.Set(1, 0, '─')
	rm_tr.Set(2, 0, '┐')
	rm_tr.Set(2, 1, '│')
	rm_tr.Set(2, 2, '│')

	rm_tl := RuneMap{[][]rune{{'.', '.', '.'}, {'.', '.', '.'}, {'.', '.', '.'}}}
	rm_tl.Set(0, 2, '│')
	rm_tl.Set(0, 1, '│')
	rm_tl.Set(0, 0, '┌')
	rm_tl.Set(1, 0, '─')
	rm_tl.Set(2, 2, '─')

	rm_br := RuneMap{[][]rune{{'.', '.', '.'}, {'.', '.', '.'}, {'.', '.', '.'}}}
	rm_br.Set(0, 2, '─')
	rm_br.Set(1, 2, '─')
	rm_br.Set(2, 2, '┘')
	rm_br.Set(2, 1, '│')
	rm_br.Set(2, 0, '│')

	rm_bl := RuneMap{[][]rune{{'.', '.', '.'}, {'.', '.', '.'}, {'.', '.', '.'}}}
	rm_bl.Set(0, 0, '│')
	rm_bl.Set(0, 1, '│')
	rm_bl.Set(0, 2, '└')
	rm_bl.Set(1, 2, '─')
	rm_bl.Set(2, 2, '─')

	type fields struct {
		Coords []int
		Type   string
	}
	tests := []struct {
		name   string
		fields fields
		want   Drawable
	}{
		{

			name: "Draw vertical line",
			fields: fields{
				Coords: []int{0, 0, 0, 2},
				Type:   "default",
			},
			want: Drawable{
				StartX: 0,
				StartY: 0,
				Content: RuneMap{[][]rune{
					{'│'},
					{'│'},
					{'│'},
				},
				},
			},
		},
		{
			name: "Draw horizontal line",
			fields: fields{
				Coords: []int{0, 0, 2, 0},
				Type:   "default",
			},
			want: Drawable{
				StartX: 0,
				StartY: 0,
				Content: RuneMap{[][]rune{
					{'─', '─', '─'},
				}},
			},
		},
		{
			name: "Draw horizontal line (reverse)",
			fields: fields{
				Coords: []int{2, 0, 0, 0},
				Type:   "default",
			},
			want: Drawable{
				StartX: 0,
				StartY: 0,
				Content: RuneMap{[][]rune{
					{'─', '─', '─'},
				}},
			},
		},
		{
			name: "Draw two lines (right-up)",
			fields: fields{
				Coords: []int{0, 2, 2, 2, 0, 2},
				Type:   "default",
			},
			want: Drawable{StartX: 0, StartY: 0, Content: rm_br},
		},
		{
			name: "Draw two lines (right-down)",
			fields: fields{
				Coords: []int{0, 0, 2, 0, 2, 2},
				Type:   "default",
			},
			want: Drawable{StartX: 0, StartY: 0, Content: rm_tr},
		},
		{
			name: "Draw two lines (left-up)",
			fields: fields{
				Coords: []int{2, 2, 0, 2, 0, 0},
				Type:   "default",
			},
			want: Drawable{StartX: 0, StartY: 0, Content: rm_bl},
		},
		{
			name: "Draw two lines (left-down)",
			fields: fields{
				Coords: []int{2, 0, 0, 0, 0, 2},
				Type:   "default",
			},
			want: Drawable{StartX: 0, StartY: 0, Content: rm_tl},
		},
		{
			name: "Draw two lines (up-left)",
			fields: fields{
				Coords: []int{2, 2, 2, 0, 0, 0},
				Type:   "default",
			},
			want: Drawable{StartX: 0, StartY: 0, Content: rm_tr},
		},
		{
			name: "Draw two lines (down-left)",
			fields: fields{
				Coords: []int{2, 0, 2, 2, 0, 2},
				Type:   "default",
			},
			want: Drawable{StartX: 0, StartY: 0, Content: rm_br},
		},
		{
			name: "Draw two lines (up-right)",
			fields: fields{
				Coords: []int{0, 2, 0, 2, 2, 0},
				Type:   "default",
			},
			want: Drawable{StartX: 0, StartY: 0, Content: rm_tl},
		},
		{
			name: "Draw two lines (down-right)",
			fields: fields{
				Coords: []int{0, 0, 0, 2, 2, 2},
				Type:   "default",
			},
			want: Drawable{StartX: 0, StartY: 0, Content: rm_bl},
		},

		{
			name: "Draw two lines, offset (right-up)",
			fields: fields{
				Coords: []int{5, 7, 8, 7, 8, 3},
				Type:   "default",
			},
			want: Drawable{StartX: 5, StartY: 3, Content: rm_br},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Line{
				Coords: tt.fields.Coords,
				Type:   tt.fields.Type,
			}
			if got := l.Drawable(); !reflect.DeepEqual(got, tt.want) {
				t.Logf("Drew:\n--START--\n%v\n--END--\n", got.Content)
				t.Logf("Expexted:\n--START--\n%v\n--END--\n", tt.want.Content)

				t.Errorf("Line.Drawable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMinMaxCoords(t *testing.T) {
	tests := []struct {
		name     string
		c        []int
		wantMinX int
		wantMinY int
		wantMaxX int
		wantMaxY int
	}{
		{
			name:     "big to small",
			c:        []int{6, 5, 4, 3, 2, 1},
			wantMinX: 2,
			wantMinY: 1,
			wantMaxX: 6,
			wantMaxY: 5,
		},

		{
			name:     "big to small",
			c:        []int{1, 2, 3, 4, 5, 6},
			wantMinX: 1,
			wantMinY: 2,
			wantMaxX: 5,
			wantMaxY: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMinX, gotMinY, gotMaxX, gotMaxY := findMinMaxCoords(tt.c)
			if gotMinX != tt.wantMinX {
				t.Errorf("findMinMaxCoords() gotMinX = %v, want %v", gotMinX, tt.wantMinX)
			}
			if gotMinY != tt.wantMinY {
				t.Errorf("findMinMaxCoords() gotMinY = %v, want %v", gotMinY, tt.wantMinY)
			}
			if gotMaxX != tt.wantMaxX {
				t.Errorf("findMinMaxCoords() gotMaxX = %v, want %v", gotMaxX, tt.wantMaxX)
			}
			if gotMaxY != tt.wantMaxY {
				t.Errorf("findMinMaxCoords() gotMaxY = %v, want %v", gotMaxY, tt.wantMaxY)
			}
		})
	}
}
