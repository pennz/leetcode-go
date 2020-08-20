package main

import "testing"

func Test_minTime(t *testing.T) {
	type args struct {
		n        int
		edges    [][]int
		hasApple []bool
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
			if got := minTime(tt.args.n, tt.args.edges, tt.args.hasApple); got != tt.want {
				t.Errorf("minTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
