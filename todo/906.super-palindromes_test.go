package main

import "testing"

func Test_superpalindromesInRange(t *testing.T) {
	type args struct {
		L string
		R string
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
			if got := superpalindromesInRange(tt.args.L, tt.args.R); got != tt.want {
				t.Errorf("superpalindromesInRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
