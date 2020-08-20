package main

import "testing"

func Test_videoStitching(t *testing.T) {
	type args struct {
		clips [][]int
		T     int
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
			if got := videoStitching(tt.args.clips, tt.args.T); got != tt.want {
				t.Errorf("videoStitching() = %v, want %v", got, tt.want)
			}
		})
	}
}
