package main

import "testing"

func Test_bagOfTokensScore(t *testing.T) {
	type args struct {
		tokens []int
		P      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bagOfTokensScore(tt.args.tokens, tt.args.P); got != tt.want {
				t.Errorf("bagOfTokensScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
