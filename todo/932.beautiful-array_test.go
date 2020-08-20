Generated Test_beautifulArray
package main

import (
	"reflect"
	"testing"
)

func Test_beautifulArray(t *testing.T) {
	type args struct {
		N int
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
			if got := beautifulArray(tt.args.N); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("beautifulArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
