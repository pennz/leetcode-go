Generated TestConstructor
Generated TestSolution_Pick
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		rects [][]int
	}
	tests := []struct {
		name string
		args args
		want Solution
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.rects); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution_Pick(t *testing.T) {
	tests := []struct {
		name string
		this *Solution
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Solution{}
			if got := this.Pick(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution.Pick() = %v, want %v", got, tt.want)
			}
		})
	}
}
