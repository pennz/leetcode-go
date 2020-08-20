// Generated TestPeekingIterator_hasNext
// Generated TestPeekingIterator_next
// Generated TestPeekingIterator_peek
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		iter *Iterator
	}
	tests := []struct {
		name string
		args args
		want *PeekingIterator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.iter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeekingIterator_hasNext(t *testing.T) {
	tests := []struct {
		name string
		this *PeekingIterator
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &PeekingIterator{}
			if got := this.hasNext(); got != tt.want {
				t.Errorf("PeekingIterator.hasNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeekingIterator_next(t *testing.T) {
	tests := []struct {
		name string
		this *PeekingIterator
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &PeekingIterator{}
			if got := this.next(); got != tt.want {
				t.Errorf("PeekingIterator.next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPeekingIterator_peek(t *testing.T) {
	tests := []struct {
		name string
		this *PeekingIterator
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &PeekingIterator{}
			if got := this.peek(); got != tt.want {
				t.Errorf("PeekingIterator.peek() = %v, want %v", got, tt.want)
			}
		})
	}
}
