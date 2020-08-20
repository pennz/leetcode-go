Generated Test_rotatedDigits
package main

import "testing"

func Test_rotatedDigits(t *testing.T) {
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
			if got := rotatedDigits(tt.args.N); got != tt.want {
				t.Errorf("rotatedDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
