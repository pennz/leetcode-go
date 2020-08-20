Generated Test_kthFactor
package main

import "testing"

func Test_kthFactor(t *testing.T) {
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
			if got := kthFactor(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("kthFactor() = %v, want %v", got, tt.want)
			}
		})
	}
}
