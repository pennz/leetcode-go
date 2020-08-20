// Generated TestSolution_GetRandom
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		head *ListNode
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
			if got := Constructor(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolution_GetRandom(t *testing.T) {
	tests := []struct {
		name string
		this *Solution
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Solution{}
			if got := this.GetRandom(); got != tt.want {
				t.Errorf("Solution.GetRandom() = %v, want %v", got, tt.want)
			}
		})
	}
}
