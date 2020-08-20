Generated Test_numSubarrayBoundedMax
package main

import "testing"

func Test_numSubarrayBoundedMax(t *testing.T) {
	type args struct {
		A []int
		L int
		R int
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
			if got := numSubarrayBoundedMax(tt.args.A, tt.args.L, tt.args.R); got != tt.want {
				t.Errorf("numSubarrayBoundedMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
