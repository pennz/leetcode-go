Generated Test_sortItems
package main

import (
	"reflect"
	"testing"
)

func Test_sortItems(t *testing.T) {
	type args struct {
		n           int
		m           int
		group       []int
		beforeItems [][]int
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
			if got := sortItems(tt.args.n, tt.args.m, tt.args.group, tt.args.beforeItems); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortItems() = %v, want %v", got, tt.want)
			}
		})
	}
}
