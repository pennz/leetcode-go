package main

import "testing"

func Test_findMinStep(t *testing.T) {
	type args struct {
		board string
		hand  string
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
			if got := findMinStep(tt.args.board, tt.args.hand); got != tt.want {
				t.Errorf("findMinStep() = %v, want %v", got, tt.want)
			}
		})
	}
}
