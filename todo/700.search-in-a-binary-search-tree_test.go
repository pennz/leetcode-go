package main

import (
	"reflect"
	"testing"
)

func Test_searchBST(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
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
			if got := searchBST(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
