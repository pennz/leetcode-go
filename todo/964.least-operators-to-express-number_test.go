package main

import "testing"

func Test_leastOpsExpressTarget(t *testing.T) {
	type args struct {
		x      int
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
			if got := leastOpsExpressTarget(tt.args.x, tt.args.target); got != tt.want {
				t.Errorf("leastOpsExpressTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
