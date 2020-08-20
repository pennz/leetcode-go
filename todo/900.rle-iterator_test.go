// Generated TestRLEIterator_Next
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want RLEIterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRLEIterator_Next(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		this *RLEIterator
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &RLEIterator{}
			if got := this.Next(tt.args.n); got != tt.want {
				t.Errorf("RLEIterator.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}
