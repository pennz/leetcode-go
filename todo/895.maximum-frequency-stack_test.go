Generated TestConstructor
Generated TestFreqStack_Push
Generated TestFreqStack_Pop
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want FreqStack
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

func TestFreqStack_Push(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		this *FreqStack
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &FreqStack{}
			this.Push(tt.args.x)
		})
	}
}

func TestFreqStack_Pop(t *testing.T) {
	tests := []struct {
		name string
		this *FreqStack
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &FreqStack{}
			if got := this.Pop(); got != tt.want {
				t.Errorf("FreqStack.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
