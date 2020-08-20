package main

import "testing"

func Test_maxSumAfterPartitioning(t *testing.T) {
	type args struct {
		A []int
		K int
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
			if got := maxSumAfterPartitioning(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("maxSumAfterPartitioning() = %v, want %v", got, tt.want)
			}
		})
	}
}
