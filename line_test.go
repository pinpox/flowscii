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


	rm_vert := RuneMap{}
	rm_vert.Set(0, 0, '│')
	rm_vert.Set(0, 1, '│')
	rm_vert.Set(0, 2, '│')


	rm_horiz := RuneMap{}
	rm_horiz.Set(0, 0, '─')
	rm_horiz.Set(1, 0, '─')
	rm_horiz.Set(2, 0, '─')


	rm_tr := RuneMap{}
	rm_tr.Set(0, 0, '─')
	rm_tr.Set(1, 0, '─')
	rm_tr.Set(2, 0, '┐')
	rm_tr.Set(2, 1, '│')
	rm_tr.Set(2, 2, '│')

	rm_tl := RuneMap{}
	rm_tl.Set(0, 2, '│')
	rm_tl.Set(0, 1, '│')
	rm_tl.Set(0, 0, '┌')
	rm_tl.Set(1, 0, '─')
	rm_tl.Set(2, 0, '─')

	rm_br := RuneMap{}
	rm_br.Set(0, 2, '─')
	rm_br.Set(1, 2, '─')
	rm_br.Set(2, 2, '┘')
	rm_br.Set(2, 1, '│')
	rm_br.Set(2, 0, '│')

	rm_br_offset := RuneMap{}
	rm_br_offset.Set(5, 7, '─')
	rm_br_offset.Set(6, 7, '─')
	rm_br_offset.Set(7, 7, '┘')
	rm_br_offset.Set(7, 6, '│')
	rm_br_offset.Set(7, 5, '│')

	rm_bl := RuneMap{}
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
		want   RuneMap
	}{
		{

			name: "Draw vertical line",
			fields: fields{
				Coords: []int{0, 0, 0, 2},
				Type:   "default",
			},
			want: rm_vert,
		},
		{
			name: "Draw horizontal line",
			fields: fields{
				Coords: []int{0, 0, 2, 0},
				Type:   "default",
			},
			want: rm_horiz,
		},
		{
			name: "Draw horizontal line (reverse)",
			fields: fields{
				Coords: []int{2, 0, 0, 0},
				Type:   "default",
			},
			want: rm_horiz,
		},
		{
			name: "Draw two lines (right-up)",
			fields: fields{
				Coords: []int{0, 2, 2, 2, 2, 0},
				Type:   "default",
			},
			want: rm_br,
		},
		{
			name: "Draw two lines (right-down)",
			fields: fields{
				Coords: []int{0, 0, 2, 0, 2, 2},
				Type:   "default",
			},
			want:  rm_tr,
		},
		{
			name: "Draw two lines (left-up)",
			fields: fields{
				Coords: []int{2, 2, 0, 2, 0, 0},
				Type:   "default",
			},
			want: rm_bl,
		},
		{
			name: "Draw two lines (left-down)",
			fields: fields{
				Coords: []int{2, 0, 0, 0, 0, 2},
				Type:   "default",
			},
			want: rm_tl,
		},
		{
			name: "Draw two lines (up-left)",
			fields: fields{
				Coords: []int{2, 2, 2, 0, 0, 0},
				Type:   "default",
			},
			want: rm_tr,
		},
		{
			name: "Draw two lines (down-left)",
			fields: fields{
				Coords: []int{2, 0, 2, 2, 0, 2},
				Type:   "default",
			},
			want: rm_br,
		},
		{
			name: "Draw two lines (up-right)",
			fields: fields{
				Coords: []int{0, 2, 0, 0, 2, 0},
				Type:   "default",
			},
			want: rm_tl,
		},
		{
			name: "Draw two lines (down-right)",
			fields: fields{
				Coords: []int{0, 0, 0, 2, 2, 2},
				Type:   "default",
			},
			want: rm_bl,
		},

		{
			name: "Draw two lines, offset (right-up)",
			fields: fields{
				Coords: []int{5, 7, 7, 7, 7, 5 },
				Type:   "default",
			},
			want: rm_br_offset,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Line{
				Coords: tt.fields.Coords,
				Type:   tt.fields.Type,
			}
			if got := l.Draw(); !reflect.DeepEqual(got, tt.want) {
				t.Logf("Drew:\n--START--\n%v\n--END--\n", got)
				t.Logf("Expexted:\n--START--\n%v\n--END--\n", tt.want)

				t.Errorf("Line.Drawable() =\n--START--\n%v\n--END--\n, want\n--START--\n%v\n--END--", got, tt.want)
			}
		})
	}
}
