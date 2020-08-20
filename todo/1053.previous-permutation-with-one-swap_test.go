Generated Test_prevPermOpt1
package main

import (
	"reflect"
	"testing"
)

func Test_prevPermOpt1(t *testing.T) {
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
			if got := prevPermOpt1(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prevPermOpt1() = %v, want %v", got, tt.want)
			}
		})
	}
}
