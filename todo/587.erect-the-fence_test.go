Generated Test_outerTrees
package main

import (
	"reflect"
	"testing"
)

func Test_outerTrees(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := outerTrees(tt.args.points); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("outerTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
