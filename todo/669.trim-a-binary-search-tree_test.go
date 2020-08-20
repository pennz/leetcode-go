package main

import (
	"reflect"
	"testing"
)

func Test_trimBST(t *testing.T) {
	type args struct {
		root *TreeNode
		L    int
		R    int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimBST(tt.args.root, tt.args.L, tt.args.R); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("trimBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
