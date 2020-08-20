package main

import "testing"

func Test_longestZigZag(t *testing.T) {
	type args struct {
		root *TreeNode
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
			if got := longestZigZag(tt.args.root); got != tt.want {
				t.Errorf("longestZigZag() = %v, want %v", got, tt.want)
			}
		})
	}
}
