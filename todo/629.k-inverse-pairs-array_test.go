package main

import "testing"

func Test_kInversePairs(t *testing.T) {
	type args struct {
		n int
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
			if got := kInversePairs(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("kInversePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
