package main

import "testing"

func Test_numberWays(t *testing.T) {
	type args struct {
		hats [][]int
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
			if got := numberWays(tt.args.hats); got != tt.want {
				t.Errorf("numberWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
