package main

import "testing"

func TestRuneMap_String(t *testing.T) {
	tests := []struct {
		name string
		rm   RuneMap
		want string
	}{
		{
			name: "Vertical line",
			rm: RuneMap{[][]rune{
				{'.', '|', '.', '.'},
				{'.', '|', '.', '.'},
				{'.', '|', '.', '.'},
			}},
			want: ".|..\n.|..\n.|..",
		},
		{
			name: "Horizontal line",
			rm: RuneMap{[][]rune{
				{'.', '.', '.', '.'},
				{'.', '.', '.', '.'},
				{'-', '-', '-', '-'},
			},
			},
			want: "----\n....\n....",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rm.String(); got != tt.want {
				t.Errorf("RuneMap.String() = \n%v\n, want \n%v", got, tt.want)
			}
		})
	}
}
