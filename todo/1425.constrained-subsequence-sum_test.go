package main

import "testing"

func Test_constrainedSubsetSum(t *testing.T) {
	type args struct {
		nums []int
		k    int
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
			if got := constrainedSubsetSum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("constrainedSubsetSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
