Generated Test_isRationalEqual
package main

import "testing"

func Test_isRationalEqual(t *testing.T) {
	type args struct {
		S string
		T string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isRationalEqual(tt.args.S, tt.args.T); got != tt.want {
				t.Errorf("isRationalEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
