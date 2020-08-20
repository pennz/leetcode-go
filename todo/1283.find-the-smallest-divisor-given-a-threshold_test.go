package main

import "testing"

func Test_smallestDivisor(t *testing.T) {
	type args struct {
		nums      []int
		threshold int
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
			if got := smallestDivisor(tt.args.nums, tt.args.threshold); got != tt.want {
				t.Errorf("smallestDivisor() = %v, want %v", got, tt.want)
			}
		})
	}
}
