package main

import "testing"

func Test_mergeStones(t *testing.T) {
	type args struct {
		stones []int
		K      int
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
			if got := mergeStones(tt.args.stones, tt.args.K); got != tt.want {
				t.Errorf("mergeStones() = %v, want %v", got, tt.want)
			}
		})
	}
}
