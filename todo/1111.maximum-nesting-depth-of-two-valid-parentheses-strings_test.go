Generated Test_maxDepthAfterSplit
package main

import (
	"reflect"
	"testing"
)

func Test_maxDepthAfterSplit(t *testing.T) {
	type args struct {
		seq string
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
			if got := maxDepthAfterSplit(tt.args.seq); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxDepthAfterSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
