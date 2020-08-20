package main

import "testing"

func Test_reachableNodes(t *testing.T) {
	type args struct {
		edges [][]int
		M     int
		N     int
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
			if got := reachableNodes(tt.args.edges, tt.args.M, tt.args.N); got != tt.want {
				t.Errorf("reachableNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
