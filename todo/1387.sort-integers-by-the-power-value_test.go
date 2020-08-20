Generated Test_getKth
package main

import "testing"

func Test_getKth(t *testing.T) {
	type args struct {
		lo int
		hi int
		k  int
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
			if got := getKth(tt.args.lo, tt.args.hi, tt.args.k); got != tt.want {
				t.Errorf("getKth() = %v, want %v", got, tt.want)
			}
		})
	}
}
