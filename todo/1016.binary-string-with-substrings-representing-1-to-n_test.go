Generated Test_queryString
package main

import "testing"

func Test_queryString(t *testing.T) {
	type args struct {
		S string
		N int
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
			if got := queryString(tt.args.S, tt.args.N); got != tt.want {
				t.Errorf("queryString() = %v, want %v", got, tt.want)
			}
		})
	}
}
