Generated Test_findGoodStrings
package main

import "testing"

func Test_findGoodStrings(t *testing.T) {
	type args struct {
		n    int
		s1   string
		s2   string
		evil string
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
			if got := findGoodStrings(tt.args.n, tt.args.s1, tt.args.s2, tt.args.evil); got != tt.want {
				t.Errorf("findGoodStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
