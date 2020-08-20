package main

import "testing"

func Test_findMaxConsecutiveOnes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name     string
		args     args
		wantCons int
	}{
		// TODO: Add test cases.
		{"basic", args{[]int{1, 1, 1, 0, 1, 1}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCons := findMaxConsecutiveOnes(tt.args.nums); gotCons != tt.wantCons {
				t.Errorf("findMaxConsecutiveOnes() = %v, want %v", gotCons, tt.wantCons)
			}
		})
	}
}
