Generated Test_maxNumberOfFamilies
package main

import "testing"

func Test_maxNumberOfFamilies(t *testing.T) {
	type args struct {
		n             int
		reservedSeats [][]int
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
			if got := maxNumberOfFamilies(tt.args.n, tt.args.reservedSeats); got != tt.want {
				t.Errorf("maxNumberOfFamilies() = %v, want %v", got, tt.want)
			}
		})
	}
}
