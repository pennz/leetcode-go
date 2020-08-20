// Generated TestRangeModule_AddRange
// Generated TestRangeModule_QueryRange
// Generated TestRangeModule_RemoveRange
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want RangeModule
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRangeModule_AddRange(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name string
		this *RangeModule
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &RangeModule{}
			this.AddRange(tt.args.left, tt.args.right)
		})
	}
}

func TestRangeModule_QueryRange(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name string
		this *RangeModule
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &RangeModule{}
			if got := this.QueryRange(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("RangeModule.QueryRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRangeModule_RemoveRange(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name string
		this *RangeModule
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &RangeModule{}
			this.RemoveRange(tt.args.left, tt.args.right)
		})
	}
}
