package main

import "testing"

func Test_isSelfCrossing(t *testing.T) {
	type args struct {
		x []int
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
			if got := isSelfCrossing(tt.args.x); got != tt.want {
				t.Errorf("isSelfCrossing() = %v, want %v", got, tt.want)
			}
		})
	}
}
