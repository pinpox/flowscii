package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestRuneMap_Get(t *testing.T) {

	rm1 := RuneMap{
		data: map[int]map[int]rune{
			0: {0: 'z'},
			1: {0: 'y'},
			2: {0: 'x'},
		},
	}

	rm2 := RuneMap{}
	rm2.Set(5, 5, 'x')

	type args struct {
		x int
		y int
	}
	tests := []struct {
		name    string
		args    args
		want    rune
		runemap RuneMap
	}{
		{
			name:    "Get from nil RuneMap",
			args:    args{0, 0},
			want:    CHAR_EMPTY,
			runemap: RuneMap{},
		},
		{
			name:    "Get from nil column",
			want:    CHAR_EMPTY,
			args:    args{7, 0},
			runemap: rm1,
		},
		{
			name:    "Get missing key on existing column",
			want:    CHAR_EMPTY,
			args:    args{1, 7},
			runemap: rm1,
		},
		{
			name:    "Get existing key on existing column",
			want:    'x',
			args:    args{2, 0},
			runemap: rm1,
		},

		{
			name:    "Get existing key on existing column (from set)",
			want:    'x',
			args:    args{5, 5},
			runemap: rm2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rm := tt.runemap
			if got := rm.Get(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("RuneMap.Get() = '%v', want '%v'", string(got), string(tt.want))
			}
		})
	}
}

func TestRuneMap_Set(t *testing.T) {
	// type fields struct {
	//	data map[int]map[int]rune
	// }

	rm1_before := RuneMap{}
	rm1_after := RuneMap{data: map[int]map[int]rune{
		3: {5: 'x'},
	}}

	rm2_before := RuneMap{
		data: map[int]map[int]rune{
			3: {
				6: 'a',
				7: 'd',
				8: CHAR_EMPTY,
			},
			4: {
				6: 'b',
				7: CHAR_EMPTY,
				8: 'e',
			},

			5: {
				6: 'c',
				7: CHAR_EMPTY,
				8: 'f',
			},
		},
	}
	rm2_after := RuneMap{


		data: map[int]map[int]rune{
			3: {
				6: 'a',
				7: 'd',
				8: CHAR_EMPTY,
			},
			4: {
				6: 'b',
				7: CHAR_EMPTY,
				8: 'e',
			},

			5: {
				6: 'x',
				7: CHAR_EMPTY,
				8: 'f',
			},
		},

	}

	// rm2.Set(3, 6, 'a')
	// rm2.Set(3, 7, 'd')
	// rm2.Set(3, 8, CHAR_EMPTY)

	// rm2.Set(4, 6, 'b')
	// rm2.Set(4, 7, CHAR_EMPTY)
	// rm2.Set(4, 8, 'e')

	// rm2.Set(5, 6, 'c')
	// rm2.Set(5, 7, CHAR_EMPTY)
	// rm2.Set(5, 8, 'f')

	type args struct {
		x int
		y int
		r rune
	}
	tests := []struct {
		name string
		rm   RuneMap
		want RuneMap
		args args
	}{

		{
			name: "Set on empty map",
			rm:   rm1_before,
			want: rm1_after,
			args: args{3, 5, 'x'},
		},
		{
			name: "Set on existing value",
			rm:   rm2_before,
			want: rm2_after,
			args: args{5, 6, 'x'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.rm.Set(tt.args.x, tt.args.y, tt.args.r)

			if !reflect.DeepEqual(tt.rm, tt.want) {
				// t.Errorf("Box.Drawable() = %v, want %v", got, tt.want)
				t.Errorf("RuneMap after Set() = \n%+v\n, want = \n%+v", tt.rm, tt.want)
			}
		})
	}
}

func TestRuneMap_MinMax(t *testing.T) {
	type fields struct {
		data map[int]map[int]rune
	}

	rm1 := RuneMap{}
	rm2 := RuneMap{}
	rm2.Set(3, 5, 'x')
	rm2.Set(6, 8, 'x')

	rm3 := RuneMap{}
	rm3.Set(6, 8, 'x')
	rm3.Set(3, 5, 'x')

	rm4 := RuneMap{}
	rm4.Set(3, 6, 'a')
	rm4.Set(3, 7, 'd')
	rm4.Set(3, 8, CHAR_EMPTY)

	rm4.Set(4, 6, 'b')
	rm4.Set(4, 7, CHAR_EMPTY)
	rm4.Set(4, 8, 'e')

	rm4.Set(5, 6, 'c')
	rm4.Set(5, 7, CHAR_EMPTY)
	rm4.Set(5, 8, 'f')

	tests := []struct {
		name   string
		rm     RuneMap
		wantX1 int
		wantY1 int
		wantX2 int
		wantY2 int
	}{

		{"Empty map", rm1, 0, 0, 0, 0},
		{"Map with entry", rm2, 3, 5, 6, 8},
		{"Map with entry (different order)", rm3, 3, 5, 6, 8},
		{"Semi-filled map", rm4, 3, 6, 5, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX1, gotY1, gotX2, gotY2 := tt.rm.MinMax()
			if gotX1 != tt.wantX1 {
				t.Errorf("RuneMap.MinMax() gotX1 = %v, want %v", gotX1, tt.wantX1)
			}
			if gotY1 != tt.wantY1 {
				t.Errorf("RuneMap.MinMax() gotY1 = %v, want %v", gotY1, tt.wantY1)
			}
			if gotX2 != tt.wantX2 {
				t.Errorf("RuneMap.MinMax() gotX2 = %v, want %v", gotX2, tt.wantX2)
			}
			if gotY2 != tt.wantY2 {
				t.Errorf("RuneMap.MinMax() gotY2 = %v, want %v", gotY2, tt.wantY2)
			}
		})
	}
}

func TestRuneMap_String(t *testing.T) {

	rm1 := RuneMap{}

	/*
		3 4 5
		6 a b c
		7 d
		8   e f
	*/

	rm2 := RuneMap{}
	rm2.Set(3, 6, 'a')
	rm2.Set(3, 7, 'd')
	rm2.Set(3, 8, CHAR_EMPTY)

	rm2.Set(4, 6, 'b')
	rm2.Set(4, 7, CHAR_EMPTY)
	rm2.Set(4, 8, 'e')

	rm2.Set(5, 6, 'c')
	rm2.Set(5, 7, CHAR_EMPTY)
	rm2.Set(5, 8, 'f')

	tests := []struct {
		name string
		rm   RuneMap
		want string
	}{
		{
			name: "Empty RuneMap",
			rm:   rm1,
			want: "",
		},

		{
			name: "Filled RuneMap",
			rm:   rm2,
			want: "abc\nd  \n ef",
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rm.String(); got != tt.want {
				t.Errorf("RuneMap.String() = \n---START\n%v\n---END\n, want \n---START\n%v\n---END\n",
					strings.ReplaceAll(got, " ", "."),
					strings.ReplaceAll(tt.want, " ", "."))
			}
		})
	}
}
