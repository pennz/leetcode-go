Generated Test_numMatchingSubseq
package main

import "testing"

func Test_numMatchingSubseq(t *testing.T) {
	type args struct {
		S     string
		words []string
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
			if got := numMatchingSubseq(tt.args.S, tt.args.words); got != tt.want {
				t.Errorf("numMatchingSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}
