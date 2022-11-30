package main

import (
	"reflect"
	"testing"
)

func TestBox_Drawable(t *testing.T) {

	rm_9x9 := RuneMap{[][]rune{{'.', '.', '.'}, {'.', '.', '.'}, {'.', '.', '.'}}}

	rm_9x9.Set(0, 2, '─')
	rm_9x9.Set(1, 2, '─')
	rm_9x9.Set(2, 2, '┐')
	rm_9x9.Set(2, 1, '│')
	rm_9x9.Set(2, 0, '│')

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
			name: "Draw 9x9 box",
			fields: fields{
				Coords: []int{0, 0, 2, 2},
				Type:   "default",
			},
			want: Drawable{StartX: 0, StartY: 0, Content: rm_9x9},
		},
		{
			name: "Draw 9x9 box (reverse)",
			fields: fields{
				Coords: []int{2, 2, 0, 0},
				Type:   "default",
			},
			want: Drawable{StartX: 0, StartY: 0, Content: rm_9x9},
		},
		{
			name: "Draw 9x9 box (other)",
			fields: fields{
				Coords: []int{0, 2, 2, 0},
				Type:   "default",
			},
			want: Drawable{StartX: 0, StartY: 0, Content: rm_9x9},
		},
		{
			name: "Draw 9x9 box (other-reverse)",
			fields: fields{
				Coords: []int{2, 0, 0, 2},
				Type:   "default",
			},
			want: Drawable{StartX: 0, StartY: 0, Content: rm_9x9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Box{
				Coords: tt.fields.Coords,
				Type:   tt.fields.Type,
			}
			if got := b.Drawable(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Box.Drawable() = %v, want %v", got, tt.want)
			}
		})
	}
}
