Generated TestConstructor
Generated TestNestedIterator_Next
Generated TestNestedIterator_HasNext
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		nestedList []*NestedInteger
	}
	tests := []struct {
		name string
		args args
		want *NestedIterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.nestedList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNestedIterator_Next(t *testing.T) {
	tests := []struct {
		name string
		this *NestedIterator
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &NestedIterator{}
			if got := this.Next(); got != tt.want {
				t.Errorf("NestedIterator.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNestedIterator_HasNext(t *testing.T) {
	tests := []struct {
		name string
		this *NestedIterator
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &NestedIterator{}
			if got := this.HasNext(); got != tt.want {
				t.Errorf("NestedIterator.HasNext() = %v, want %v", got, tt.want)
			}
		})
	}
}
