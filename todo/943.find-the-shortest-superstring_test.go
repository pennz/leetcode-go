package main

import "testing"

func Test_shortestSuperstring(t *testing.T) {
	type args struct {
		A []string
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
			if got := shortestSuperstring(tt.args.A); got != tt.want {
				t.Errorf("shortestSuperstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
