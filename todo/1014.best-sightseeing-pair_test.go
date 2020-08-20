package main

import "testing"

func Test_maxScoreSightseeingPair(t *testing.T) {
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
			if got := maxScoreSightseeingPair(tt.args.A); got != tt.want {
				t.Errorf("maxScoreSightseeingPair() = %v, want %v", got, tt.want)
			}
		})
	}
}
