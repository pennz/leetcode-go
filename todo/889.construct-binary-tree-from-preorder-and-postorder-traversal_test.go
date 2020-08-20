package main

import (
	"reflect"
	"testing"
)

func Test_constructFromPrePost(t *testing.T) {
	type args struct {
		pre  []int
		post []int
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
			if got := constructFromPrePost(tt.args.pre, tt.args.post); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructFromPrePost() = %v, want %v", got, tt.want)
			}
		})
	}
}
