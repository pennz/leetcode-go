package main

import "testing"

func Test_partitionDisjoint(t *testing.T) {
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
			if got := partitionDisjoint(tt.args.A); got != tt.want {
				t.Errorf("partitionDisjoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
