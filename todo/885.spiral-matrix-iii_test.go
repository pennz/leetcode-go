package main

import (
	"reflect"
	"testing"
)

func Test_spiralMatrixIII(t *testing.T) {
	type args struct {
		R  int
		C  int
		r0 int
		c0 int
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
			if got := spiralMatrixIII(tt.args.R, tt.args.C, tt.args.r0, tt.args.c0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralMatrixIII() = %v, want %v", got, tt.want)
			}
		})
	}
}
