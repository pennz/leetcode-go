Generated Test_sumZero
package main

import (
	"reflect"
	"testing"
)

func Test_sumZero(t *testing.T) {
	type args struct {
		n int
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
			if got := sumZero(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
