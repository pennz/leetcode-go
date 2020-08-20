package main

import "testing"

func Test_brokenCalc(t *testing.T) {
	type args struct {
		X int
		Y int
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
			if got := brokenCalc(tt.args.X, tt.args.Y); got != tt.want {
				t.Errorf("brokenCalc() = %v, want %v", got, tt.want)
			}
		})
	}
}
