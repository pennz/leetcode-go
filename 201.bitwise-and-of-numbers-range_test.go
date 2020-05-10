package main

import (
	"testing"
)

func Test_rangeBitwiseAnd(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"😤", args{3, 7}, 0}, // e.g. 101, 110, 111, 100 and -> 100, so we just test half range, if we have for then we got 0
		{"hard", args{0, 1}, 0},
		{"hard2", args{0, 8}, 0},
		{"hard3", args{1, 8}, 0},
		{"my_error", args{6, 7}, 6},       // 110, 111, only the last bit changed
		{"🚩my_error2!!", args{1, 3}, 0},   // 01, 10, 11, -> 0 (my algorithm: get 01, but it can change and masked,so we
		{"🚩my_error2-try", args{2, 6}, 0}, // 01, 10, 11, -> 0 (my algorithm: get 01, but it can change and masked,so we
		// another example to make it more clear, 010 ~ 110 -> 0
		{"corner", args{0, 0}, 0},
		{"easy2", args{5, 7}, 4}, // e.g. 101, 110, 111, and -> 100, so we just test half range, if we have for then we got 0
		{"easy2", args{4, 7}, 4}, // e.g. 101, 110, 111, 100 and -> 100, so we just test half range, if we have for then we got 0
		{"easy2", args{7, 7}, 7}, // e.g. 101, 110, 111, and -> 100, so we just test half range, if we have for then we got 0
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.want = rangeBitwiseAndForce(tt.args.m, tt.args.n)
			if got := rangeBitwiseAnd(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("rangeBitwiseAnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getHighestBitsValue(t *testing.T) {
	type args struct {
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{9}, 8},
		{"", args{255}, 255},
		{"", args{253}, 252},
		{"", args{128}, 128},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHighestBitsValue(tt.args.m); got != tt.want {
				t.Errorf("getHighestBitValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rangeBitwiseAndForce(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"corner", args{0, 0}, 0},
		{"easy", args{0, 1}, 0},
		{"easy2", args{5, 7}, 4}, // e.g. 101, 110, 111, and -> 100, so we just test half range, if we have for then we got 0
		{"easy2", args{0, 8}, 0},
		{"easy2", args{1, 8}, 0},
		{"easy2", args{4, 7}, 4}, // e.g. 101, 110, 111, and -> 100, so we just test half range, if we have for then we got 0
		{"easy2", args{7, 7}, 7}, // e.g. 101, 110, 111, and -> 100, so we just test half range, if we have for then we got 0
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeBitwiseAndForce(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("rangeBitwiseAndForce() = %v, want %v", got, tt.want)
			}
		})
	}
}
