package main

import "testing"

func Test_tilingRectangle(t *testing.T) {
	type args struct {
		n int
		m int
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
			if got := tilingRectangle(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("tilingRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
