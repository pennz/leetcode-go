Generated Test_networkDelayTime
package main

import "testing"

func Test_networkDelayTime(t *testing.T) {
	type args struct {
		times [][]int
		N     int
		K     int
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
			if got := networkDelayTime(tt.args.times, tt.args.N, tt.args.K); got != tt.want {
				t.Errorf("networkDelayTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
