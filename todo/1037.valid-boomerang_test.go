Generated Test_isBoomerang
package main

import "testing"

func Test_isBoomerang(t *testing.T) {
	type args struct {
		points [][]int
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
			if got := isBoomerang(tt.args.points); got != tt.want {
				t.Errorf("isBoomerang() = %v, want %v", got, tt.want)
			}
		})
	}
}