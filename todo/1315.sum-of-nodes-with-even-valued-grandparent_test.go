package main

import "testing"

func Test_sumEvenGrandparent(t *testing.T) {
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
			if got := sumEvenGrandparent(tt.args.root); got != tt.want {
				t.Errorf("sumEvenGrandparent() = %v, want %v", got, tt.want)
			}
		})
	}
}
