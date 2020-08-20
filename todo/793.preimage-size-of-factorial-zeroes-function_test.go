package main

import "testing"

func Test_preimageSizeFZF(t *testing.T) {
	type args struct {
		K int
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
			if got := preimageSizeFZF(tt.args.K); got != tt.want {
				t.Errorf("preimageSizeFZF() = %v, want %v", got, tt.want)
			}
		})
	}
}
