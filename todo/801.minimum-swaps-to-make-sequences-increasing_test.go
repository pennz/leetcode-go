package main

import "testing"

func Test_minSwap(t *testing.T) {
	type args struct {
		A []int
		B []int
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
			if got := minSwap(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("minSwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
