package main

import (
	"reflect"
	"testing"
)

func Test_sumEvenAfterQueries(t *testing.T) {
	type args struct {
		A       []int
		queries [][]int
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
			if got := sumEvenAfterQueries(tt.args.A, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumEvenAfterQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
