Generated Test_findInMountainArray
package main

import "testing"

func Test_findInMountainArray(t *testing.T) {
	type args struct {
		target      int
		mountainArr *MountainArray
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
			if got := findInMountainArray(tt.args.target, tt.args.mountainArr); got != tt.want {
				t.Errorf("findInMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
