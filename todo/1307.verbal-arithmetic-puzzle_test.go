package main

import "testing"

func Test_isSolvable(t *testing.T) {
	type args struct {
		words  []string
		result string
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
			if got := isSolvable(tt.args.words, tt.args.result); got != tt.want {
				t.Errorf("isSolvable() = %v, want %v", got, tt.want)
			}
		})
	}
}
