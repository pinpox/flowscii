package main

import (
	"reflect"
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
			args: args{ x: 0, y: 2 },
			want: 'a',
		},

		{
			name: "top_right",
			rm: RuneMap{[][]rune{
				{'e', 'f', 'g', 'h'},
				{'a', 'b', 'c', 'd'},
			} },
			args: args{ x: 3, y: 1 },
			want: 'd',
		},

		{
			name: "bottom_left",
			rm: RuneMap{[][]rune{
				{'i', 'j', 'k', 'l'},
				{'e', 'f', 'g', 'h'},
				{'a', 'b', 'c', 'd'},
			} },
			args: args{ x: 1, y: 1 },
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
		want  int
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.rm.Dims()
			if got != tt.want {
				t.Errorf("RuneMap.Dims() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("RuneMap.Dims() got1 = %v, want %v", got1, tt.want1)
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
		name string
		args args
		want RuneMap
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initRuneMap(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initRuneMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
