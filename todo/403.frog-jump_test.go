package main

import "testing"

func Test_canCross(t *testing.T) {
	type args struct {
		stones []int
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
			if got := canCross(tt.args.stones); got != tt.want {
				t.Errorf("canCross() = %v, want %v", got, tt.want)
			}
		})
	}
}
