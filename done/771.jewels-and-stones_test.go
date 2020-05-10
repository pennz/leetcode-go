package main

import (
	"testing"
)

func Test_numJewelsInStones(t *testing.T) {
	type args struct {
		J string
		S string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"stone_jew",
			args{"aA", "aaAAAbbbc"},
			5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numJewelsInStones(tt.args.J, tt.args.S); got != tt.want {
				t.Errorf("numJewelsInStones() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numJewelsInStones2(t *testing.T) {
	type args struct {
		J string
		S string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"stone_jew",
			args{"aA", "aaAAAbbbc"},
			5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numJewelsInStones2(tt.args.J, tt.args.S); got != tt.want {
				t.Errorf("numJewelsInStones2() = %v, want %v", got, tt.want)
			}
		})
	}
}
