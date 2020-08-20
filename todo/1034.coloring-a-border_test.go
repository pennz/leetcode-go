package main

import (
	"reflect"
	"testing"
)

func Test_colorBorder(t *testing.T) {
	type args struct {
		grid  [][]int
		r0    int
		c0    int
		color int
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
			if got := colorBorder(tt.args.grid, tt.args.r0, tt.args.c0, tt.args.color); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("colorBorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
