package main

import "testing"

func Test_minAreaRect(t *testing.T) {
	type args struct {
		points [][]int
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
			if got := minAreaRect(tt.args.points); got != tt.want {
				t.Errorf("minAreaRect() = %v, want %v", got, tt.want)
			}
		})
	}
}
