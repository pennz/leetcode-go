package main

import (
	"reflect"
	"testing"
)

func Test_shortestAlternatingPaths(t *testing.T) {
	type args struct {
		n          int
		red_edges  [][]int
		blue_edges [][]int
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
			if got := shortestAlternatingPaths(tt.args.n, tt.args.red_edges, tt.args.blue_edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortestAlternatingPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
