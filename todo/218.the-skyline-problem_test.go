package main

import (
	"reflect"
	"testing"
)

func Test_getSkyline(t *testing.T) {
	type args struct {
		buildings [][]int
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
			if got := getSkyline(tt.args.buildings); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSkyline() = %v, want %v", got, tt.want)
			}
		})
	}
}
