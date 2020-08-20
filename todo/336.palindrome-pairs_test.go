package main

import (
	"reflect"
	"testing"
)

func Test_palindromePairs(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := palindromePairs(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("palindromePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
