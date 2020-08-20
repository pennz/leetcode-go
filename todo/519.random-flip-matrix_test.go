Generated TestConstructor
Generated TestSolution_Flip
Generated TestSolution_Reset
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		n_rows int
		n_cols int
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
			if got := Constructor(tt.args.n_rows, tt.args.n_cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution_Flip(t *testing.T) {
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
			if got := this.Flip(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution.Flip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution_Reset(t *testing.T) {
	tests := []struct {
		name string
		this *Solution
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Solution{}
			this.Reset()
		})
	}
}
