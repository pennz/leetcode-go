package main

import "testing"

func Test_maskPII(t *testing.T) {
	type args struct {
		S string
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
			if got := maskPII(tt.args.S); got != tt.want {
				t.Errorf("maskPII() = %v, want %v", got, tt.want)
			}
		})
	}
}
