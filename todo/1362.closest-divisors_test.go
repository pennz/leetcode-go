Generated Test_closestDivisors
package main

import (
	"reflect"
	"testing"
)

func Test_closestDivisors(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closestDivisors(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("closestDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}
