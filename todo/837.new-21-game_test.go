package main

import "testing"

func Test_new21Game(t *testing.T) {
	type args struct {
		N int
		K int
		W int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := new21Game(tt.args.N, tt.args.K, tt.args.W); got != tt.want {
				t.Errorf("new21Game() = %v, want %v", got, tt.want)
			}
		})
	}
}
