Generated Test_angleClock
package main

import "testing"

func Test_angleClock(t *testing.T) {
	type args struct {
		hour    int
		minutes int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := angleClock(tt.args.hour, tt.args.minutes); got != tt.want {
				t.Errorf("angleClock() = %v, want %v", got, tt.want)
			}
		})
	}
}
