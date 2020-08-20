// Generated TestSolution_Pick
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

func TestSolution_Pick(t *testing.T) {
	type args struct {
		target int
	}
	tests := []struct {
		name string
		this *Solution
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Solution{}
			if got := this.Pick(tt.args.target); got != tt.want {
				t.Errorf("Solution.Pick() = %v, want %v", got, tt.want)
			}
		})
	}
}
