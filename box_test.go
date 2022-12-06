package main

import (
	"reflect"
	"testing"
)

func TestBox_Drawable(t *testing.T) {

	rm_9x9 := RuneMap{}
	rm_9x9.Set(0, 2, '└')
	rm_9x9.Set(2, 2, '┘')
	rm_9x9.Set(2, 0, '┐')
	rm_9x9.Set(0, 0, '┌')

	rm_9x9.Set(1, 0, '─')
	rm_9x9.Set(1, 2, '─')
	rm_9x9.Set(0, 1, '│')
	rm_9x9.Set(2, 1, '│')

	rm_9x9_shadow := RuneMap{}
	rm_9x9_shadow.Set(1, 3, '└')
	rm_9x9_shadow.Set(3, 3, '┘')
	rm_9x9_shadow.Set(3, 1, '┐')
	rm_9x9_shadow.Set(1, 1, '┌')
	rm_9x9_shadow.Set(2, 1, '─')
	rm_9x9_shadow.Set(2, 3, '─')
	rm_9x9_shadow.Set(1, 2, '│')
	rm_9x9_shadow.Set(3, 2, '│')
	rm_9x9_shadow.Set(4, 2, '░')
	rm_9x9_shadow.Set(4, 3, '░')
	rm_9x9_shadow.Set(4, 4, '░')
	rm_9x9_shadow.Set(3, 4, '░')
	rm_9x9_shadow.Set(2, 4, '░')


	rm_9x9_shadow_offset := RuneMap{}
	rm_9x9_shadow_offset.Set(2, 5, '└')
	rm_9x9_shadow_offset.Set(4, 5, '┘')
	rm_9x9_shadow_offset.Set(4, 3, '┐')
	rm_9x9_shadow_offset.Set(2, 3, '┌')
	rm_9x9_shadow_offset.Set(3, 3, '─')
	rm_9x9_shadow_offset.Set(3, 5, '─')
	rm_9x9_shadow_offset.Set(2, 4, '│')
	rm_9x9_shadow_offset.Set(4, 4, '│')
	rm_9x9_shadow_offset.Set(5, 4, '░')
	rm_9x9_shadow_offset.Set(5, 5, '░')
	rm_9x9_shadow_offset.Set(5, 6, '░')
	rm_9x9_shadow_offset.Set(4, 6, '░')
	rm_9x9_shadow_offset.Set(3, 6, '░')

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
			name:   "Draw 9x9 box, default",
			fields: fields{Coords: []int{0, 0, 2, 2}, Type: "default"},
			want:   rm_9x9,
		},
		// {
		// 	name:   "Draw 9x9 box, default (reverse)",
		// 	fields: fields{Coords: []int{2, 2, 0, 0}, Type: "default"},
		// 	want:   rm_9x9,
		// },
		// {
		// 	name:   "Draw 9x9 box, default (other)",
		// 	fields: fields{Coords: []int{0, 2, 2, 0}, Type: "default"},
		// 	want:   rm_9x9,
		// },
		// {
		// 	name:   "Draw 9x9 box, default (other-reverse)",
		// 	fields: fields{Coords: []int{2, 0, 0, 2}, Type: "default"},
		// 	want:   rm_9x9,
		// },
		// {
		// 	name:   "Draw 9x9 box, default, offset",
		// 	fields: fields{Coords: []int{2, 1, 4, 3}, Type: "default"},
		// 	want:   rm_9x9,
		// },
		{
			name:   "Draw 9x9 box, shadow",
			fields: fields{Coords: []int{1, 1, 3, 3}, Type: "shadow"},
			want:   rm_9x9_shadow,
		},
		// {
		// 	name:   "Draw 9x9 box, shadow (reverse)",
		// 	fields: fields{Coords: []int{2, 2, 0, 0}, Type: "shadow"},
		// 	want:   rm_9x9_shadow,
		// },
		// {
		// 	name:   "Draw 9x9 box, shadow (other)",
		// 	fields: fields{Coords: []int{0, 2, 2, 0}, Type: "shadow"},
		// 	want:   rm_9x9_shadow,
		// },
		// {
		// 	name:   "Draw 9x9 box, shadow (other-reverse)",
		// 	fields: fields{Coords: []int{2, 0, 0, 2}, Type: "shadow"},
		// 	want:   rm_9x9_shadow,
		// },
		{
			name:   "Draw 9x9 box, shadow, offset",
			fields: fields{Coords: []int{2, 3, 4, 5}, Type: "shadow"},
			want:   rm_9x9_shadow_offset,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Box{
				Coords: tt.fields.Coords,
				Type:   tt.fields.Type,
			}
			if got := b.Draw(); !reflect.DeepEqual(got, tt.want) {
				// t.Errorf("Box.Drawable() = %+v, want %+v", got, tt.want)
				t.Errorf("Box.Drawable() =\n--START--\n%v\n--END--, want\n--START--\n%v\n--END--\nCoords got\n%v\nCoords want\n\v%v", got, tt.want, got.data, tt.want.data)
			}
		})
	}
}
