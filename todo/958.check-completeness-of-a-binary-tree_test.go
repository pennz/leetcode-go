package main

import "testing"

func Test_isCompleteTree(t *testing.T) {
	type args struct {
		root *TreeNode
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
			if got := isCompleteTree(tt.args.root); got != tt.want {
				t.Errorf("isCompleteTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
