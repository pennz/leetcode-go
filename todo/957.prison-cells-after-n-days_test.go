package main

import (
	"reflect"
	"testing"
)

func Test_prisonAfterNDays(t *testing.T) {
	type args struct {
		cells []int
		N     int
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
			if got := prisonAfterNDays(tt.args.cells, tt.args.N); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prisonAfterNDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
