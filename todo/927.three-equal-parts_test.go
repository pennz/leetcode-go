Generated Test_threeEqualParts
package main

import (
	"reflect"
	"testing"
)

func Test_threeEqualParts(t *testing.T) {
	type args struct {
		A []int
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
			if got := threeEqualParts(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeEqualParts() = %v, want %v", got, tt.want)
			}
		})
	}
}
