package main

import "testing"

func Test_minSumOfLengths(t *testing.T) {
	type args struct {
		arr    []int
		target int
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
			if got := minSumOfLengths(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("minSumOfLengths() = %v, want %v", got, tt.want)
			}
		})
	}
}
