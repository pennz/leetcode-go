Generated Test_countOrders
package main

import "testing"

func Test_countOrders(t *testing.T) {
	type args struct {
		n int
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
			if got := countOrders(tt.args.n); got != tt.want {
				t.Errorf("countOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}
