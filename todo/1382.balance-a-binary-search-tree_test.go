package main

import (
	"reflect"
	"testing"
)

func Test_balanceBST(t *testing.T) {
	type args struct {
		root *TreeNode
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
			if got := balanceBST(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("balanceBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
