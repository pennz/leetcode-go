package main

import (
	"reflect"
	"testing"
)

func Test_findDuplicate(t *testing.T) {
	type args struct {
		paths []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate(tt.args.paths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
