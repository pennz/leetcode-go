Generated Test_ways
package main

import "testing"

func Test_ways(t *testing.T) {
	type args struct {
		pizza []string
		k     int
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
			if got := ways(tt.args.pizza, tt.args.k); got != tt.want {
				t.Errorf("ways() = %v, want %v", got, tt.want)
			}
		})
	}
}
