Generated Test_strongPasswordChecker
package main

import "testing"

func Test_strongPasswordChecker(t *testing.T) {
	type args struct {
		s string
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
			if got := strongPasswordChecker(tt.args.s); got != tt.want {
				t.Errorf("strongPasswordChecker() = %v, want %v", got, tt.want)
			}
		})
	}
}
