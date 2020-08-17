package main

import "testing"

func Test_isNStraightHand(t *testing.T) {
	type args struct {
		hand []int
		W    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"basic", args{[]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3}, true},
		{"basic2", args{[]int{1, 2, 3, 6, 2, 3, 4, 7, 9}, 3}, false},
		{"basic3", args{[]int{1, 2, 3, 6, 2, 3, 4, 7}, 4}, false},
		{"corner1", args{[]int{1, 2, 3}, 4}, false},
		{"corner2", args{[]int{1, 2, 3}, 3}, true},
		{"corner2", args{[]int{1, 2, 3}, 1}, true},
		{"corner2", args{[]int{1, 2, 3}, 2}, false},
		{"failed2", args{[]int{1, 1, 2, 2, 3, 3}, 2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNStraightHand(tt.args.hand, tt.args.W); got != tt.want {
				t.Errorf("isNStraightHand() = %v, want %v", got, tt.want)
			}
		})
	}
}
