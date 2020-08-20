package main

import "testing"

func Test_minNumberOfSemesters(t *testing.T) {
	type args struct {
		n            int
		dependencies [][]int
		k            int
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
			if got := minNumberOfSemesters(tt.args.n, tt.args.dependencies, tt.args.k); got != tt.want {
				t.Errorf("minNumberOfSemesters() = %v, want %v", got, tt.want)
			}
		})
	}
}
