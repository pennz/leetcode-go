Generated TestConstructor
Generated TestDinnerPlates_Push
Generated TestDinnerPlates_Pop
Generated TestDinnerPlates_PopAtStack
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name string
		args args
		want DinnerPlates
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.capacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDinnerPlates_Push(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		this *DinnerPlates
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &DinnerPlates{}
			this.Push(tt.args.val)
		})
	}
}

func TestDinnerPlates_Pop(t *testing.T) {
	tests := []struct {
		name string
		this *DinnerPlates
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &DinnerPlates{}
			if got := this.Pop(); got != tt.want {
				t.Errorf("DinnerPlates.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDinnerPlates_PopAtStack(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name string
		this *DinnerPlates
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &DinnerPlates{}
			if got := this.PopAtStack(tt.args.index); got != tt.want {
				t.Errorf("DinnerPlates.PopAtStack() = %v, want %v", got, tt.want)
			}
		})
	}
}
