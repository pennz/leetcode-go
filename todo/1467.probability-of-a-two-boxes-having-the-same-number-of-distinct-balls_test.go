package main

import "testing"

func Test_getProbability(t *testing.T) {
	type args struct {
		balls []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getProbability(tt.args.balls); got != tt.want {
				t.Errorf("getProbability() = %v, want %v", got, tt.want)
			}
		})
	}
}
