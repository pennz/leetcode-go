Generated Test_numDupDigitsAtMostN
package main

import "testing"

func Test_numDupDigitsAtMostN(t *testing.T) {
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
			if got := numDupDigitsAtMostN(tt.args.N); got != tt.want {
				t.Errorf("numDupDigitsAtMostN() = %v, want %v", got, tt.want)
			}
		})
	}
}
