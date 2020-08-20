package main

import "testing"

func Test_splitArraySameAverage(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitArraySameAverage(tt.args.A); got != tt.want {
				t.Errorf("splitArraySameAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}
