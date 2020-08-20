Generated Test_numPoints
package main

import "testing"

func Test_numPoints(t *testing.T) {
	type args struct {
		points [][]int
		r      int
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
			if got := numPoints(tt.args.points, tt.args.r); got != tt.want {
				t.Errorf("numPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
