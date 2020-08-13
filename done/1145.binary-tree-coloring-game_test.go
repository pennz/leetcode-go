package main

import "testing"

func Test_btreeGameWinningMove(t *testing.T) {
	type args struct {
		root *TreeNode
		n    int
		x    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"basic", args{&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}, 3, 3}, true},
		{"basic_2", args{&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}, 3, 1}, false},
		{"basic_3", args{&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}, 3, 2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := btreeGameWinningMove(tt.args.root, tt.args.n, tt.args.x); got != tt.want {
				t.Errorf("btreeGameWinningMove() = %v, want %v", got, tt.want)
			}
		})
	}
}
