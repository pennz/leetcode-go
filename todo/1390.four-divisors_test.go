Generated Test_sumFourDivisors
package main

import "testing"

func Test_sumFourDivisors(t *testing.T) {
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
			if got := sumFourDivisors(tt.args.nums); got != tt.want {
				t.Errorf("sumFourDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}
