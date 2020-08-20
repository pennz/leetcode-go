package main

import "testing"

func Test_checkRecord(t *testing.T) {
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
			if got := checkRecord(tt.args.n); got != tt.want {
				t.Errorf("checkRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}
