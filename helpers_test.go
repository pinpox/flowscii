package main

import (
	"testing"
)

func TestRuneMap_Get(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		rm   RuneMap
		args args
		want rune
	}{
		{
			name: "top_left",
			rm: RuneMap{[][]rune{
				{'i', 'j', 'k', 'l'},
				{'e', 'f', 'g', 'h'},
				{'a', 'b', 'c', 'd'},
			}},
			args: args{x: 0, y: 2},
			want: 'a',
		},

		{
			name: "top_right",
			rm: RuneMap{[][]rune{
				{'e', 'f', 'g', 'h'},
				{'a', 'b', 'c', 'd'},
			}},
			args: args{x: 3, y: 1},
			want: 'd',
		},

		{
			name: "bottom_left",
			rm: RuneMap{[][]rune{
				{'i', 'j', 'k', 'l'},
				{'e', 'f', 'g', 'h'},
				{'a', 'b', 'c', 'd'},
			}},
			args: args{x: 1, y: 1},
			want: 'f',
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rm.Get(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("RuneMap.Get() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestRuneMap_Set(t *testing.T) {
	type args struct {
		x int
		y int
		r rune
	}
	tests := []struct {
		name string
		rm   RuneMap
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.rm.Set(tt.args.x, tt.args.y, tt.args.r)
		})
	}
}

func TestRuneMap_Dims(t *testing.T) {
	tests := []struct {
		name  string
		rm    RuneMap
		wantX int
		wantY int
	}{
		{
			name: "9x9",
			rm: RuneMap{
				[][]rune{
					{'.', '.', '.'},
					{'.', '.', '.'},
					{'.', '.', '.'},
				},
			},
			wantX: 3,
			wantY: 3,
		},

		{
			name: "5x2",
			rm: RuneMap{
				[][]rune{
					{'.', '.', '.', '.', '.'},
					{'.', '.', '.', '.', '.'},
				},
			},
			wantX: 5,
			wantY: 2,
		},

		{
			name: "2x5",
			rm: RuneMap{
				[][]rune{
					{'.', '.'},
					{'.', '.'},
					{'.', '.'},
					{'.', '.'},
					{'.', '.'},
				},
			},
			wantX: 2,
			wantY: 5,
		},

		{
			name: "1x3",
			rm: RuneMap{
				[][]rune{
					{'.'},
					{'.'},
					{'.'},
				},
			},
			wantX: 1,
			wantY: 3,
		},
		// panics
		// {
		// 	name: "0x0",
		// 	rm: RuneMap{
		// 		[][]rune{},
		// 	},
		// 	wantX: 0,
		// 	wantY: 0,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := tt.rm.Dims()
			if gotX != tt.wantX {
				t.Errorf("RuneMap.Dims() gotX = %v, wantX %v", gotX, tt.wantX)
			}
			if gotY != tt.wantY {
				t.Errorf("RuneMap.Dims() gotY = %v, wantY %v", gotY, tt.wantY)
			}
		})
	}
}

func Test_initRuneMap(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name     string
		args     args
		wantDimX int
		wantDimY int
	}{
		{
			name:     "3x3",
			args:     args{3, 3},
			wantDimX: 3,
			wantDimY: 3,
		},
		{
			name:     "1x3",
			args:     args{1, 3},
			wantDimX: 1,
			wantDimY: 3,
		},
		{
			name:     "3x1",
			args:     args{3, 1},
			wantDimX: 3,
			wantDimY: 1,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := initRuneMap(tt.args.x, tt.args.y).Dims()
			if gotX != tt.wantDimX {
				t.Errorf("initRuneMap.Dims() gotX = %v, wantDimX %v", gotX, tt.wantDimX)
			}
			if gotY != tt.wantDimY {
				t.Errorf("initRuneMap.Dims() gotY = %v, wantDimY %v", gotY, tt.wantDimY)
			}
		})
	}
}
