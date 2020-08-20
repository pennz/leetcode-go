package main

import "testing"

func Test_lenLongestFibSubseq(t *testing.T) {
	type args struct {
		A []int
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
			if got := lenLongestFibSubseq(tt.args.A); got != tt.want {
				t.Errorf("lenLongestFibSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}
