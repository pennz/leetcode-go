package main

import "testing"

func Test_findMaximumXOR(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"b", args{[]int{2, 3}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaximumXOR(tt.args.nums); got != tt.want {
				t.Errorf("findMaximumXOR() = %v, want %v", got, tt.want)
			}
		})
	}
}
