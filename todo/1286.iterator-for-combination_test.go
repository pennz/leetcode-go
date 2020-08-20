// Generated TestCombinationIterator_Next
// Generated TestCombinationIterator_HasNext
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		characters        string
		combinationLength int
	}
	tests := []struct {
		name string
		args args
		want CombinationIterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.characters, tt.args.combinationLength); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCombinationIterator_Next(t *testing.T) {
	tests := []struct {
		name string
		this *CombinationIterator
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &CombinationIterator{}
			if got := this.Next(); got != tt.want {
				t.Errorf("CombinationIterator.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCombinationIterator_HasNext(t *testing.T) {
	tests := []struct {
		name string
		this *CombinationIterator
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &CombinationIterator{}
			if got := this.HasNext(); got != tt.want {
				t.Errorf("CombinationIterator.HasNext() = %v, want %v", got, tt.want)
			}
		})
	}
}
