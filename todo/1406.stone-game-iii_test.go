package main

import "testing"

func Test_stoneGameIII(t *testing.T) {
	type args struct {
		stoneValue []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stoneGameIII(tt.args.stoneValue); got != tt.want {
				t.Errorf("stoneGameIII() = %v, want %v", got, tt.want)
			}
		})
	}
}
