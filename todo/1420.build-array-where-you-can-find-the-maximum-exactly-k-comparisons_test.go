Generated Test_numOfArrays
package main

import "testing"

func Test_numOfArrays(t *testing.T) {
	type args struct {
		n int
		m int
		k int
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
			if got := numOfArrays(tt.args.n, tt.args.m, tt.args.k); got != tt.want {
				t.Errorf("numOfArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
