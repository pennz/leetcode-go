package main

import "testing"

func Test_checkOverlap(t *testing.T) {
	type args struct {
		radius   int
		x_center int
		y_center int
		x1       int
		y1       int
		x2       int
		y2       int
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
			if got := checkOverlap(tt.args.radius, tt.args.x_center, tt.args.y_center, tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2); got != tt.want {
				t.Errorf("checkOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}
