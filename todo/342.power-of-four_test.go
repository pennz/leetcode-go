Generated Test_isPowerOfFour
package main

import "testing"

func Test_isPowerOfFour(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfFour(tt.args.num); got != tt.want {
				t.Errorf("isPowerOfFour() = %v, want %v", got, tt.want)
			}
		})
	}
}