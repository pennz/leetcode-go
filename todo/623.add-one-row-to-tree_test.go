package main

import (
	"reflect"
	"testing"
)

func Test_addOneRow(t *testing.T) {
	type args struct {
		root *TreeNode
		v    int
		d    int
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
			if got := addOneRow(tt.args.root, tt.args.v, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addOneRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
