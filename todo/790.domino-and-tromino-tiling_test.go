package main

import "testing"

func Test_numTilings(t *testing.T) {
	type args struct {
		N int
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
			if got := numTilings(tt.args.N); got != tt.want {
				t.Errorf("numTilings() = %v, want %v", got, tt.want)
			}
		})
	}
}
