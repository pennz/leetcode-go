Generated Test_countPrimeSetBits
package main

import "testing"

func Test_countPrimeSetBits(t *testing.T) {
	type args struct {
		L int
		R int
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
			if got := countPrimeSetBits(tt.args.L, tt.args.R); got != tt.want {
				t.Errorf("countPrimeSetBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
