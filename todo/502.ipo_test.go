package main

import "testing"

func Test_findMaximizedCapital(t *testing.T) {
	type args struct {
		k       int
		W       int
		Profits []int
		Capital []int
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
			if got := findMaximizedCapital(tt.args.k, tt.args.W, tt.args.Profits, tt.args.Capital); got != tt.want {
				t.Errorf("findMaximizedCapital() = %v, want %v", got, tt.want)
			}
		})
	}
}
