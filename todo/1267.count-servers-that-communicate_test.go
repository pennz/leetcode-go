package main

import "testing"

func Test_countServers(t *testing.T) {
	type args struct {
		grid [][]int
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
			if got := countServers(tt.args.grid); got != tt.want {
				t.Errorf("countServers() = %v, want %v", got, tt.want)
			}
		})
	}
}
