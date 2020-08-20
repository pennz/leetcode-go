// Generated TestCustomStack_Push
// Generated TestCustomStack_Pop
// Generated TestCustomStack_Increment
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		maxSize int
	}
	tests := []struct {
		name string
		args args
		want CustomStack
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.maxSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomStack_Push(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		this *CustomStack
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &CustomStack{}
			this.Push(tt.args.x)
		})
	}
}

func TestCustomStack_Pop(t *testing.T) {
	tests := []struct {
		name string
		this *CustomStack
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &CustomStack{}
			if got := this.Pop(); got != tt.want {
				t.Errorf("CustomStack.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomStack_Increment(t *testing.T) {
	type args struct {
		k   int
		val int
	}
	tests := []struct {
		name string
		this *CustomStack
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &CustomStack{}
			this.Increment(tt.args.k, tt.args.val)
		})
	}
}
