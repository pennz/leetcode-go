package main

import "testing"

func Test_countPalindromicSubsequences(t *testing.T) {
	type args struct {
		S string
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
			if got := countPalindromicSubsequences(tt.args.S); got != tt.want {
				t.Errorf("countPalindromicSubsequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
