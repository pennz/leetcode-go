Generated Test_prefixesDivBy5
package main

import (
	"reflect"
	"testing"
)

func Test_prefixesDivBy5(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prefixesDivBy5(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prefixesDivBy5() = %v, want %v", got, tt.want)
			}
		})
	}
}
