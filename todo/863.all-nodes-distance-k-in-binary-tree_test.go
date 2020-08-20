package main

import (
	"reflect"
	"testing"
)

func Test_distanceK(t *testing.T) {
	type args struct {
		root   *TreeNode
		target *TreeNode
		K      int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distanceK(tt.args.root, tt.args.target, tt.args.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("distanceK() = %v, want %v", got, tt.want)
			}
		})
	}
}
