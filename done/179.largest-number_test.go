package main

import "testing"

func Test_largestNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"zeros", args{[]int{0, 0, 0}}, "0"},     // 121 12, compare 12-> same, then crop [12] continue, 12 -> 1, so we choose 12
		{"same", args{[]int{12, 12}}, "1212"},    // 121 12, compare 12-> same, then crop [12] continue, 12 -> 1, so we choose 12
		{"hard2", args{[]int{121, 12}}, "12121"}, // 121 12, compare 12-> same, then crop [12] continue, 12 -> 1, so we choose 12
		{"easy", args{[]int{9, 1}}, "91"},
		{"medium", args{[]int{51, 52}}, "5251"},
		{"hard", args{[]int{3, 30, 34, 5, 9}}, "9534330"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestNumber(tt.args.nums); got != tt.want {
				t.Errorf("largestNumber() = %v, want %v, given input %v", got, tt.want, tt.args)
			}
		})
	}
}
