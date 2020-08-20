Generated Test_canBeEqual
package main

import "testing"

func Test_canBeEqual(t *testing.T) {
	type args struct {
		target []int
		arr    []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canBeEqual(tt.args.target, tt.args.arr); got != tt.want {
				t.Errorf("canBeEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
