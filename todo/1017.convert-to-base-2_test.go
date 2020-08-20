package main

import "testing"

func Test_baseNeg2(t *testing.T) {
	type args struct {
		N int
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
			if got := baseNeg2(tt.args.N); got != tt.want {
				t.Errorf("baseNeg2() = %v, want %v", got, tt.want)
			}
		})
	}
}
