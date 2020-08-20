package main

import "testing"

func Test_removeBoxes(t *testing.T) {
	type args struct {
		boxes []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeBoxes(tt.args.boxes); got != tt.want {
				t.Errorf("removeBoxes() = %v, want %v", got, tt.want)
			}
		})
	}
}
