Generated Test_numSubseq
package main

import "testing"

func Test_numSubseq(t *testing.T) {
	type args struct {
		nums   []int
		target int
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
			if got := numSubseq(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("numSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}
