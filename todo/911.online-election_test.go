Generated TestConstructor
Generated TestTopVotedCandidate_Q
package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		persons []int
		times   []int
	}
	tests := []struct {
		name string
		args args
		want TopVotedCandidate
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.persons, tt.args.times); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTopVotedCandidate_Q(t *testing.T) {
	type args struct {
		t int
	}
	tests := []struct {
		name string
		this *TopVotedCandidate
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &TopVotedCandidate{}
			if got := this.Q(tt.args.t); got != tt.want {
				t.Errorf("TopVotedCandidate.Q() = %v, want %v", got, tt.want)
			}
		})
	}
}
