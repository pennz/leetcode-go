package main

import "testing"

func Test_soupServings(t *testing.T) {
	type args struct {
		N int
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
			if got := soupServings(tt.args.N); got != tt.want {
				t.Errorf("soupServings() = %v, want %v", got, tt.want)
			}
		})
	}
}
