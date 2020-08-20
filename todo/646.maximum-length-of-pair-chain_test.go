Generated Test_findLongestChain
package main

import "testing"

func Test_findLongestChain(t *testing.T) {
	type args struct {
		pairs [][]int
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
			if got := findLongestChain(tt.args.pairs); got != tt.want {
				t.Errorf("findLongestChain() = %v, want %v", got, tt.want)
			}
		})
	}
}
