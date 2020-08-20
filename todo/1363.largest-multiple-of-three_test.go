package main

import "testing"

func Test_largestMultipleOfThree(t *testing.T) {
	type args struct {
		digits []int
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
			if got := largestMultipleOfThree(tt.args.digits); got != tt.want {
				t.Errorf("largestMultipleOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
