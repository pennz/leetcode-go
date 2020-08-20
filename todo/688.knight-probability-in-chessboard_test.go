package main

import "testing"

func Test_knightProbability(t *testing.T) {
	type args struct {
		N int
		K int
		r int
		c int
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
			if got := knightProbability(tt.args.N, tt.args.K, tt.args.r, tt.args.c); got != tt.want {
				t.Errorf("knightProbability() = %v, want %v", got, tt.want)
			}
		})
	}
}
