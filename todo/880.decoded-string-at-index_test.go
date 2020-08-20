Generated Test_decodeAtIndex
package main

import "testing"

func Test_decodeAtIndex(t *testing.T) {
	type args struct {
		S string
		K int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeAtIndex(tt.args.S, tt.args.K); got != tt.want {
				t.Errorf("decodeAtIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
