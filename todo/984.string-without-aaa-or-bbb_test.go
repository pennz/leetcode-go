package main

import "testing"

func Test_strWithout3a3b(t *testing.T) {
	type args struct {
		A int
		B int
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
			if got := strWithout3a3b(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("strWithout3a3b() = %v, want %v", got, tt.want)
			}
		})
	}
}
