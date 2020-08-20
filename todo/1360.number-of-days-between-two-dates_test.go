package main

import "testing"

func Test_daysBetweenDates(t *testing.T) {
	type args struct {
		date1 string
		date2 string
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
			if got := daysBetweenDates(tt.args.date1, tt.args.date2); got != tt.want {
				t.Errorf("daysBetweenDates() = %v, want %v", got, tt.want)
			}
		})
	}
}
