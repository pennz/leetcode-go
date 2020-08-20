Generated Test_maxAbsValExpr
package main

import "testing"

func Test_maxAbsValExpr(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
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
			if got := maxAbsValExpr(tt.args.arr1, tt.args.arr2); got != tt.want {
				t.Errorf("maxAbsValExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}
