package main

import "testing"

func Test_maxChunksToSorted(t *testing.T) {
	type args struct {
		arr []int
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
			if got := maxChunksToSorted(tt.args.arr); got != tt.want {
				t.Errorf("maxChunksToSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
