package main

import "testing"

func Test_reachingPoints(t *testing.T) {
	type args struct {
		sx int
		sy int
		tx int
		ty int
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
			if got := reachingPoints(tt.args.sx, tt.args.sy, tt.args.tx, tt.args.ty); got != tt.want {
				t.Errorf("reachingPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
