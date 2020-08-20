Generated Test_circularPermutation
package main

import (
	"reflect"
	"testing"
)

func Test_circularPermutation(t *testing.T) {
	type args struct {
		n     int
		start int
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
			if got := circularPermutation(tt.args.n, tt.args.start); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("circularPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
