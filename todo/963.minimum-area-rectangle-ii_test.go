package main

import "testing"

func Test_minAreaFreeRect(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minAreaFreeRect(tt.args.points); got != tt.want {
				t.Errorf("minAreaFreeRect() = %v, want %v", got, tt.want)
			}
		})
	}
}
