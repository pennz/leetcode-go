package main

import "testing"

func Test_validateBinaryTreeNodes(t *testing.T) {
	type args struct {
		n          int
		leftChild  []int
		rightChild []int
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
			if got := validateBinaryTreeNodes(tt.args.n, tt.args.leftChild, tt.args.rightChild); got != tt.want {
				t.Errorf("validateBinaryTreeNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
