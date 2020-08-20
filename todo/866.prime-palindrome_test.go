package main

import "testing"

func Test_primePalindrome(t *testing.T) {
	type args struct {
		N int
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
			if got := primePalindrome(tt.args.N); got != tt.want {
				t.Errorf("primePalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
