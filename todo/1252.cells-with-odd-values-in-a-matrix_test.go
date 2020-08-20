package main

import "testing"

func Test_oddCells(t *testing.T) {
	type args struct {
		n       int
		m       int
		indices [][]int
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
			if got := oddCells(tt.args.n, tt.args.m, tt.args.indices); got != tt.want {
				t.Errorf("oddCells() = %v, want %v", got, tt.want)
			}
		})
	}
}
