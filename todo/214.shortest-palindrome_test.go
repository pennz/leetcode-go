Generated Test_shortestPalindrome
package main

import "testing"

func Test_shortestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("shortestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
