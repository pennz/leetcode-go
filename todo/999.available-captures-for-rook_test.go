package main

import "testing"

func Test_numRookCaptures(t *testing.T) {
	type args struct {
		board [][]byte
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
			if got := numRookCaptures(tt.args.board); got != tt.want {
				t.Errorf("numRookCaptures() = %v, want %v", got, tt.want)
			}
		})
	}
}
