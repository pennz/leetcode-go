package main

import "testing"

func Test_largestPerimeter(t *testing.T) {
	type args struct {
		A []int
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
			if got := largestPerimeter(tt.args.A); got != tt.want {
				t.Errorf("largestPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
