package main

import (
	"reflect"
	"testing"
)

func Test_deckRevealedIncreasing(t *testing.T) {
	type args struct {
		deck []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deckRevealedIncreasing(tt.args.deck); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deckRevealedIncreasing() = %v, want %v", got, tt.want)
			}
		})
	}
}
