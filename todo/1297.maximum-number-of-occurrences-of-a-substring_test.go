Generated Test_maxFreq
package main

import "testing"

func Test_maxFreq(t *testing.T) {
	type args struct {
		s          string
		maxLetters int
		minSize    int
		maxSize    int
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
			if got := maxFreq(tt.args.s, tt.args.maxLetters, tt.args.minSize, tt.args.maxSize); got != tt.want {
				t.Errorf("maxFreq() = %v, want %v", got, tt.want)
			}
		})
	}
}
