package main

import "testing"

func Test_nthMagicalNumber(t *testing.T) {
	type args struct {
		N int
		A int
		B int
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
			if got := nthMagicalNumber(tt.args.N, tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("nthMagicalNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
