Generated Test_movesToMakeZigzag
package main

import "testing"

func Test_movesToMakeZigzag(t *testing.T) {
	type args struct {
		nums []int
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
			if got := movesToMakeZigzag(tt.args.nums); got != tt.want {
				t.Errorf("movesToMakeZigzag() = %v, want %v", got, tt.want)
			}
		})
	}
}
