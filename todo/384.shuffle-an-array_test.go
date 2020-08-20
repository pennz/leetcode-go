// Generated TestSolution_Reset
// Generated TestSolution_Shuffle
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		nums []int
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
			if got := Constructor(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution_Reset(t *testing.T) {
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
			if got := this.Reset(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution.Reset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution_Shuffle(t *testing.T) {
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
			if got := this.Shuffle(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution.Shuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}
