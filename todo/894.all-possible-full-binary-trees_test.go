Generated Test_allPossibleFBT
package main

import (
	"reflect"
	"testing"
)

func Test_allPossibleFBT(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allPossibleFBT(tt.args.N); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allPossibleFBT() = %v, want %v", got, tt.want)
			}
		})
	}
}
