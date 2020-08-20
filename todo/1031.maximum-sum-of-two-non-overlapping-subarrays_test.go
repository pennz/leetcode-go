package main

import "testing"

func Test_maxSumTwoNoOverlap(t *testing.T) {
	type args struct {
		A []int
		L int
		M int
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
			if got := maxSumTwoNoOverlap(tt.args.A, tt.args.L, tt.args.M); got != tt.want {
				t.Errorf("maxSumTwoNoOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}
