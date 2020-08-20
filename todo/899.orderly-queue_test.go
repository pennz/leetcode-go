package main

import "testing"

func Test_orderlyQueue(t *testing.T) {
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
			if got := orderlyQueue(tt.args.S, tt.args.K); got != tt.want {
				t.Errorf("orderlyQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}
