package main

import "testing"

func Test_numSimilarGroups(t *testing.T) {
	type args struct {
		A []string
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
			if got := numSimilarGroups(tt.args.A); got != tt.want {
				t.Errorf("numSimilarGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
