// Generated TestMinStack_Push
// Generated TestMinStack_Pop
// Generated TestMinStack_Top
// Generated TestMinStack_GetMin
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want MinStack
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

func TestMinStack_Push(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		this *MinStack
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MinStack{}
			this.Push(tt.args.x)
		})
	}
}

func TestMinStack_Pop(t *testing.T) {
	tests := []struct {
		name string
		this *MinStack
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MinStack{}
			this.Pop()
		})
	}
}

func TestMinStack_Top(t *testing.T) {
	tests := []struct {
		name string
		this *MinStack
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MinStack{}
			if got := this.Top(); got != tt.want {
				t.Errorf("MinStack.Top() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinStack_GetMin(t *testing.T) {
	tests := []struct {
		name string
		this *MinStack
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &MinStack{}
			if got := this.GetMin(); got != tt.want {
				t.Errorf("MinStack.GetMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
