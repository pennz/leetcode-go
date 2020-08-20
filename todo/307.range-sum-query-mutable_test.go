// Generated TestNumArray_Update
// Generated TestNumArray_SumRange
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
		want NumArray
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

func TestNumArray_Update(t *testing.T) {
	type args struct {
		i   int
		val int
	}
	tests := []struct {
		name string
		this *NumArray
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &NumArray{}
			this.Update(tt.args.i, tt.args.val)
		})
	}
}

func TestNumArray_SumRange(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		this *NumArray
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &NumArray{}
			if got := this.SumRange(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("NumArray.SumRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
