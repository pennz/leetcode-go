package utils

import "testing"

func Test_quickSort(t *testing.T) {
	type args struct {
		nums        []string
		compareFunc func(string, string) bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort(tt.args.nums, tt.args.compareFunc)
		})
	}
}

func Test_partition(t *testing.T) {
	type args struct {
		nums        []string
		compareFunc func(string, string) bool
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
			if got := partition(tt.args.nums, tt.args.compareFunc); got != tt.want {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
